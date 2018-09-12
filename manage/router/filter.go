package router

import "github.com/valyala/fasthttp"

// Filter is a func for for fasthttp
type Filter func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

// BuildHandler is make a new handler with filters
func BuildHandler(baseHandler fasthttp.RequestHandler, filters ...Filter) fasthttp.RequestHandler {
	handler := baseHandler
	for _, filter := range filters {
		handler = filter(handler)
	}
	return handler
}
