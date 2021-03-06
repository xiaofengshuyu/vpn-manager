package user

import (
	"context"

	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
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
	err = h.UserService.RegisterUserWithoutCheck(context.Background(), req)
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

// Logout user logout
func (h *Handler) Logout(ctx *fasthttp.RequestCtx) {
	token := string(ctx.Request.Header.Peek("token"))
	if token == "" {
		h.WriteJSON(ctx, nil, common.NewRequestParamsValueError("token is empty"))
		return
	}
	err := h.UserService.Logout(context.Background(), token)
	h.WriteJSON(ctx, nil, err)
}

// RefreshVertifyCode refresh user's vertify code
func (h *Handler) RefreshVertifyCode(ctx *fasthttp.RequestCtx) {
	parse := gjson.ParseBytes(ctx.PostBody())
	email := parse.Get("email").String()
	err := h.UserService.EmailResend(context.Background(), &models.CommonUser{
		Email: email,
	})
	h.WriteJSON(ctx, nil, err)
}

// ResetPassword reset password
func (h *Handler) ResetPassword(ctx *fasthttp.RequestCtx) {
	req, err := userRegisterDecode(ctx)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	err = h.UserService.ResetPassword(context.Background(), req)
	h.WriteJSON(ctx, nil, err)
}
