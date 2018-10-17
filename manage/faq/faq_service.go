package faq

import (
	"context"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// Service FAQ service
type Service interface {
	GetFrequentAskedQuestions(ctx context.Context) (faqs []models.FrequentQuestions, err error)
	PushFeedBack(ctx context.Context, fb models.Feedback) (err error)
	GetFeedBack(ctx context.Context) (feedBacks []models.Feedback, err error)
}

// BaseFAQService is FAQ service implement
type BaseFAQService struct {
}

// GetFrequentAskedQuestions return enabled frequent asked questions
func (s *BaseFAQService) GetFrequentAskedQuestions(ctx context.Context) (faqs []models.FrequentQuestions, err error) {
	db := common.DB
	err = db.Where("status = ?", models.FAQStatusPublish).Find(&faqs).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}

// PushFeedBack push user feedcack content
func (s *BaseFAQService) PushFeedBack(ctx context.Context, fb models.Feedback) (err error) {
	db := common.DB
	err = db.Create(&fb).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}

// GetFeedBack get all feedback
func (s *BaseFAQService) GetFeedBack(ctx context.Context) (feedBacks []models.Feedback, err error) {
	db := common.DB
	err = db.Find(&feedBacks).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}
