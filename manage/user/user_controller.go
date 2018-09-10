package user

import (
	"context"
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	validator "gopkg.in/go-playground/validator.v9"
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
	// check
	validate := validator.New()
	err = req.Valid(validate)
	fmt.Println("=============", err)

	if err != nil {
		fmt.Println("ERR", err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("invalid validate error", err)
			h.WriteJSON(ctx, nil, err)
			return
		}
	}
	fmt.Println("mmmmmmmmmmmmm")
	h.WriteJSON(ctx, nil, nil)

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
