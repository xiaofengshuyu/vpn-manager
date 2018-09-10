package router

import (
	"encoding/base64"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
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
	internalRouter.POST("/api/manage/user", user.InsertUser)
	VPNManageRouter = RecoverWrap(BasicAuth(internalRouter.Handler, config.AppConfig.Auth.User, config.AppConfig.Auth.Password))

	// external router config
	externalRouter := fasthttprouter.New()
	// user register
	externalRouter.POST("/api/register", userHandler.Register)
	UserAccessRouter = RecoverWrap(externalRouter.Handler)
}

// basicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func basicAuth(ctx *fasthttp.RequestCtx) (username, password string, ok bool) {
	auth := ctx.Request.Header.Peek("Authorization")
	if auth == nil {
		return
	}
	return parseBasicAuth(string(auth))
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// BasicAuth is the basic auth handler
func BasicAuth(h fasthttp.RequestHandler, requiredUser, requiredPassword string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := basicAuth(ctx)
		fmt.Println(user, password, hasAuth)
		fmt.Println(requiredUser, requiredPassword)
		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(ctx)
			return
		}
		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
	})
}

// RecoverWrap is a middleware for fasthttp handler
func RecoverWrap(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		var msg string
		defer func() {
			r := recover()
			if r != nil {
				common.Logger.Error(r)
				msg = fmt.Sprintf("%v\n%s", r, debug.Stack())
				ctx.Error(msg, fasthttp.StatusInternalServerError)
			}
		}()
		h(ctx)
	})
}
