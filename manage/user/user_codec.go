package user

import (
	"encoding/json"
	"strings"

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

func userRegisterDecode(ctx *fasthttp.RequestCtx) (user *models.CommonUser, err error) {
	req := new(models.CommonUser)
	err = json.Unmarshal(ctx.PostBody(), req)
	if err != nil {
		err = common.NewRequestParamsDecodeError(err)
		return
	}
	user = &models.CommonUser{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	if user.UserName == "" {
		if index := strings.Index(user.Email, "@"); index > 0 {
			user.UserName = user.Email[:index]
		}
	}
	if validate := user.Validate(); validate != "" {
		err = common.NewRequestParamsValueError(validate)
		return
	}
	return
}
