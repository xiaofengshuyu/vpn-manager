package order

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// CommitOrderRequest is commit an order request
type CommitOrderRequest struct {
	OrderNumber string `json:"orderNumber"`
	ProductCode string `json:"productCode"`
}

func commitOrderDecode(ctx *fasthttp.RequestCtx) (order models.Order, err error) {

	err = json.Unmarshal(ctx.PostBody(), &order)
	if err != nil {
		err = common.NewRequestParamsDecodeError(err)
		return
	}
	if order.OrderNumber == "" {
		err = common.NewRequestParamsValueError()
		return
	}
	return
}
