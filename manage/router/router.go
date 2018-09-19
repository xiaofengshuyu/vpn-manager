package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/user"
)

var (
	// VPNManageRouter internal service route
	VPNManageRouter fasthttp.RequestHandler
	// UserAccessRouter external service route
	UserAccessRouter fasthttp.RequestHandler
)

func init() {
	// init user handler
	userHandler := &user.Handler{}
	userHandler.UserService = &user.BaseUserService{}

	// internal toute config
	internalRouter := fasthttprouter.New()
	VPNManageRouter = BuildHandler(
		internalRouter.Handler,
		BasicAuth,
		RecoverWrap,
	)

	// external router config
	externalRouter := fasthttprouter.New()
	// user register
	externalRouter.POST("/api/register", userHandler.Register)
	UserAccessRouter = BuildHandler(
		externalRouter.Handler,
		RecoverWrap,
	)
}
