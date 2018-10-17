package order

import (
	"context"

	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
)

// Handler is order handler
type Handler struct {
	common.BaseHandler
	OrderService Service
}

// GetProduct get product handler
func (h *Handler) GetProduct(ctx *fasthttp.RequestCtx) {
	products, err := h.OrderService.GetProduct(context.Background())
	// TODO Encode
	h.WriteJSON(ctx, products, err)
}

// CommitOrder commit an order
func (h *Handler) CommitOrder(ctx *fasthttp.RequestCtx) {
	parse := gjson.ParseBytes(ctx.PostBody())
	data := parse.Get("receipt-data").String()
	user, ok := common.GlobalSession.GetUser(ctx.ID())
	if !ok {
		h.WriteJSON(ctx, nil, common.NewNotLoginError())
		return
	}
	userCtx := context.WithValue(context.Background(), common.UserInfoKey, user)
	err := h.OrderService.CommitAnOrder(userCtx, data)
	h.WriteJSON(ctx, nil, err)
}
