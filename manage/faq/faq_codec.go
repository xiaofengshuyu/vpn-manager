package faq

import (
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// FrequentQuestionsResponse faq response
type FrequentQuestionsResponse struct {
	List []FrequentQuestionsItem `json:"list"`
}

// FrequentQuestionsItem faq item
type FrequentQuestionsItem struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func faqEncode(data []models.FrequentQuestions) (res FrequentQuestionsResponse) {
	list := make([]FrequentQuestionsItem, len(data))
	for i, item := range data {
		list[i] = FrequentQuestionsItem{
			Question: item.Question,
			Answer:   item.Answer,
		}
	}
	return FrequentQuestionsResponse{
		List: list,
	}
}

func feedbackDecode(ctx *fasthttp.RequestCtx) (fb models.Feedback, err error) {
	parse := gjson.ParseBytes(ctx.PostBody())
	fb.Question = parse.Get("question").String()
	fb.Email = parse.Get("email").String()
	fb.Phone = parse.Get("phone").String()
	// check
	if fb.Question == "" {
		err = common.NewRequestParamsDecodeError("question error")
		return
	}
	return
}
