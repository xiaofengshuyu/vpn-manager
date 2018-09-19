package lines

import (
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
)

// Handler is user handler
type Handler struct {
	common.BaseHandler
	UserService Service
}

// get lines address api
func (h *Handler) GetLines(ctx *fasthttp.RequestCtx) {
	// req, err := userRegister(ctx)
	// if err != nil {
	// 	h.WriteJSON(ctx, nil, err)
	// 	return
	// }
	// err = h.UserService.RegisterUser(context.Background(), req)
	// h.WriteJSON(ctx, nil, err)
}
