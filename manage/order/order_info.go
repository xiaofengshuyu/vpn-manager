package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// apple url
const (
	AppleSandBoxURL = "https://sandbox.itunes.apple.com/verifyReceipt"
	AppleOnlineURL  = "https://buy.itunes.apple.com/verifyReceipt "
	retry           = 3
)

// GetOrderFromApple get order information from apple server
func GetOrderFromApple(data string) (orders []*models.Order, err error) {
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
	var res *http.Response
	for i := 0; i < retry; i++ {
		res, err = http.Post(reqURL, "application/json;charset=utf-8", req)
		if err == nil {
			break
		}
	}
	if err != nil {
		logger.Error(err)
		err = common.NewAppleServerAccessError(err)
		return
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		err = common.NewAppleServerAccessError(err)
		return
	}
	logger.Debug(string(resData))

	appleOrder := AppleOrder{}
	err = json.Unmarshal(resData, &appleOrder)
	if err != nil {
		err = common.NewAppleDataInvalidError(err)
		return
	}

	// apple check
	if appleOrder.Status != 0 {
		err = common.NewAppleDataInvalidError(fmt.Sprintf("order is invalid,status is %d", appleOrder.Status))
		return
	}
	if appleOrder.Receipt.BundleID != config.AppConfig.AppleStore.BundleID {
		err = common.NewAppleDataInvalidError(fmt.Sprintf("bundle id is invalid"))
		return
	}
	if len(appleOrder.Receipt.InApp) == 0 {
		err = common.NewAppleDataInvalidError(fmt.Sprintf("in_app length is 0"))
		return
	}
	// get all product
	db := common.DB
	products := make([]*models.Product, 0)
	err = db.Model(&models.Product{}).Find(&products).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	orders = make([]*models.Order, 0)
	for _, item := range appleOrder.Receipt.InApp {
		// check order
		for _, prod := range products {
			if item.ProductID == prod.Code {
				order := &models.Order{
					OrderNumber: item.TransactionID,
					OrderData:   data,
					OrderTime:   item.PurchaseDateMs,
					ProductID:   prod.ID,
				}
				order.Quantity, _ = strconv.Atoi(item.Quantity)
				order.AddMonth = order.Quantity * prod.Duration
				orders = append(orders, order)
			}
		}
	}
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
