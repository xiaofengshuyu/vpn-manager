package faq

import (
	"context"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// Handler faq handler
type Handler struct {
	common.BaseHandler
	Service Service
}

// GetFAQ get frequent asked questions
func (h *Handler) GetFAQ(ctx *fasthttp.RequestCtx) {
	languageHeader := ctx.Request.Header.Peek("Accept-Language")
	lang := models.GetLanguage(string(languageHeader))
	data, err := h.Service.GetFrequentAskedQuestions(context.Background(), lang)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	h.WriteJSON(ctx, faqEncode(data), nil)
	return
}

// PushFeedback push feedback
func (h *Handler) PushFeedback(ctx *fasthttp.RequestCtx) {
	req, err := feedbackDecode(ctx)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	err = h.Service.PushFeedBack(context.Background(), req)
	h.WriteJSON(ctx, nil, err)
}
