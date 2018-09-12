package user

import (
	"context"
	"fmt"

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
	req, err := userRegister(ctx)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	err = h.UserService.RegisterUser(context.Background(), req)
	h.WriteJSON(ctx, nil, err)
}

// InsertUser insert user
func InsertUser(ctx *fasthttp.RequestCtx) {
	//
	fmt.Println("insert User")
}

// UpdatetUser is update user
func UpdatetUser(ctx *fasthttp.RequestCtx) {
	//
}

// GetUser is get all user
func GetUser(ctx *fasthttp.RequestCtx) {
	//
}

// DeleteUser is delete user
func DeleteUser(ctx *fasthttp.RequestCtx) {
	//
}

// GetUserByID is get one user
func GetUserByID(ctx *fasthttp.RequestCtx) {
	//
}
