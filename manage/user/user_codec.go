package user

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// RegisterRequest is user register request
type RegisterRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func userRegister(ctx *fasthttp.RequestCtx) (user *models.CommonUser, err error) {
	req := new(models.CommonUser)
	fmt.Print(req.UserName)
	err = json.Unmarshal(ctx.PostBody(), req)
	if err != nil {
		err = common.NewRequestParamsDecodeError(err)
	}
	user = &models.CommonUser{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	return
}
