package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// apple url
const (
	AppleSandBoxURL = "https://sandbox.itunes.apple.com/verifyReceipt"
	AppleOnlineURL  = ""
)

// GetOrderFromApple get order information from apple server
func GetOrderFromApple(data string) (order models.Order, err error) {
	// send request to server
	receiptData := map[string]string{
		"receipt-data": data,
	}
	req := &bytes.Buffer{}
	json.NewEncoder(req).Encode(receiptData)
	var reqURL string
	if config.AppConfig.Mode == config.PROD {
		reqURL = AppleOnlineURL
	} else {
		reqURL = AppleSandBoxURL
	}
	res, err := http.Post(reqURL, "application/json;charset=utf-8", req)
	if err != nil {
		logger.Error(err)
		return
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug(string(resData))

	appleOrder := AppleOrder{}
	err = json.Unmarshal(resData, &appleOrder)
	if err != nil {
		return
	}

	// apple check
	if appleOrder.Status != 0 {
		err = fmt.Errorf("order is invalid,status is %d", appleOrder.Status)
		return
	}
	if appleOrder.Receipt.BundleID != config.AppConfig.AppleStore.BundleID {
		err = fmt.Errorf("bundle id is invalid")
		return
	}
	if len(appleOrder.Receipt.InApp) == 0 {
		err = fmt.Errorf("in_app length is 0")
		return
	}
	// get order info
	// get last order
	lastOrder := appleOrder.Receipt.InApp[len(appleOrder.Receipt.InApp)-1]
	order.OrderData = data
	order.OrderNumber = lastOrder.TransactionID
	order.Quantity, _ = strconv.Atoi(lastOrder.Quantity)
	db := common.DB
	product := &models.Product{Code: lastOrder.ProductID}
	err = db.Where(product).First(product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = common.NewResourcesNotFoundError(fmt.Sprintf("product %s not found", lastOrder.ProductID))
		}
		return
	}
	order.Product = *product
	order.ProductID = product.ID

	return
}

// AppleOrder apple order info
type AppleOrder struct {
	Receipt     AppleReceipt `json:"receipt"`
	Status      int          `json:"status"`
	Environment string       `json:"environment"`
}

// AppleReceipt apple receipt
type AppleReceipt struct {
	ReceiptType string       `json:"receipt_type"`
	BundleID    string       `json:"bundle_id"`
	InApp       []PayHistory `json:"in_app"`
}

// PayHistory pay history
type PayHistory struct {
	Quantity       string `json:"quantity"`
	ProductID      string `json:"product_id"`
	TransactionID  string `json:"transaction_id"`
	PurchaseDateMs string `json:"purchase_date_ms"`
}
