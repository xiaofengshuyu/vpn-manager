package router

import (
	"fmt"
	"runtime/debug"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
)

// RecoverWrap is a middleware for fasthttp handler
func RecoverWrap(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		var msg string
		defer func() {
			r := recover()
			if r != nil {
				msg = fmt.Sprintf("%v\n%s", r, debug.Stack())
				common.Logger.Error(msg)
				ctx.Error(msg, fasthttp.StatusInternalServerError)
			}
		}()
		h(ctx)
	})
}
