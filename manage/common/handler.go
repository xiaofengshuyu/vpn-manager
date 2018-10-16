package common

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// BaseResponse is base response struct
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// BaseHandler is fasthttp base handler
type BaseHandler struct {
}

// WriteJSON is write data to response with json format
func (h *BaseHandler) WriteJSON(ctx *fasthttp.RequestCtx, data interface{}, err error) {
	// set head
	ctx.SetContentType("application/json;charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)

	// set body
	res := &BaseResponse{}
	if err != nil {
		res.Code, res.Message = GetErrorInfo(err)
		switch t := err.(type) {
		case *DBAccessError:
			Logger.Error(t.Stack())
		}

	} else {
		res.Data = data
	}
	d, _ := json.Marshal(res)
	// add log
	AccessLogger.Info(formatResponseData(ctx.PostBody(), d))
	ctx.SetBody(d)
}

// ResponseRecord response record
type ResponseRecord struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

func formatResponseData(req, res []byte) string {
	d, _ := json.Marshal(ResponseRecord{
		Request:  string(req),
		Response: string(res),
	})
	return string(d)
}
