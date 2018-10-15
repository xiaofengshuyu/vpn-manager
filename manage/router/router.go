package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/faq"
	"github.com/xiaofengshuyu/vpn-manager/manage/host"
	"github.com/xiaofengshuyu/vpn-manager/manage/order"
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

	// init order handler
	orderHandler := &order.Handler{}
	orderHandler.OrderService = &order.BaseOrderService{}

	// init vpn config handler
	configHandler := &host.ConfigHandler{}
	configHandler.Service = &host.BaseConfigService{}

	// init faq handler
	faqHandler := &faq.Handler{}
	faqHandler.Service = &faq.BaseFAQService{}

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
	externalRouter.POST("/api/login", userHandler.Login)

	// faq
	externalRouter.POST("/api/faq", faqHandler.GetFAQ)
	externalRouter.GET("/api/faq", faqHandler.GetFAQ)
	externalRouter.POST("/api/feedback", faqHandler.PushFeedback)

	// need login
	needLoginRouter := fasthttprouter.New()
	needLoginRouter.POST("/api/common/order/commit", orderHandler.CommitOrder)
	needLoginRouter.POST("/api/common/order/product", orderHandler.GetProduct)
	needLoginRouter.POST("/api/common/config/self", configHandler.GetVPNConfig)
	needLoginRouter.POST("/api/common/config/hosts", configHandler.GetHostList)

	externalRouter.POST("/api/common/*any", BuildHandler(
		needLoginRouter.Handler,
		VPNUserFilter,
	))

	UserAccessRouter = BuildHandler(
		externalRouter.Handler,
		RecoverWrap,
	)
}
