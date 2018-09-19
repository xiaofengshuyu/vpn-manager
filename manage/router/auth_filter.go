package router

import (
	"encoding/base64"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
	"github.com/xiaofengshuyu/vpn-manager/manage/user"
)

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
func BasicAuth(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := basicAuth(ctx)
		if hasAuth && user == config.AppConfig.Auth.User && password == config.AppConfig.Auth.Password {
			// Delegate request to the given handle
			h(ctx)
			return
		}
		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
	})
}

// VPNUserFilter is the filter for user which registered
func VPNUserFilter(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// get token
		token := ctx.Request.Header.Peek("token")
		// check token
		if ok, recorder := user.TokenCheck(string(token)); ok {
			common.GlobalSession.SetUser(ctx.ID(), recorder.User)
			h(ctx)
			common.GlobalSession.Delete(ctx.ID())
			return
		}
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	})
}
