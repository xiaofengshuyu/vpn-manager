package user

import (
	"context"

	"github.com/tidwall/gjson"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
)

// Handler is user handler
type Handler struct {
	common.BaseHandler
	UserService Service
}

// Register user register
func (h *Handler) Register(ctx *fasthttp.RequestCtx) {
	req, err := userRegisterDecode(ctx)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	err = h.UserService.RegisterUser(context.Background(), req)
	h.WriteJSON(ctx, nil, err)
}

// Login user login
func (h *Handler) Login(ctx *fasthttp.RequestCtx) {
	parse := gjson.ParseBytes(ctx.PostBody())
	username := parse.Get("username").String()
	password := parse.Get("password").String()
	if username == "" || password == "" {
		h.WriteJSON(ctx, nil, common.NewRequestParamsValueError("username or password is empty"))
		return
	}
	recorder, err := h.UserService.Login(context.Background(), username, password)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	h.WriteJSON(ctx, map[string]string{
		"token":   recorder.Token,
		"refresh": recorder.RefreshToken,
	}, nil)
}
