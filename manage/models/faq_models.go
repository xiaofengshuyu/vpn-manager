package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// FAQ status
const (
	FAQStatusDefault = iota
	FAQStatusPublish
)

// FrequentQuestions frequently asked questions
type FrequentQuestions struct {
	gorm.Model
	Question string
	Answer   string
	Status   int
	Language string
}

// feedback status
const (
	FeedbackStatusCreated = iota
	FeedbackStatusRead
	FeedbackStatusResolved
)

// Feedback user feedback
type Feedback struct {
	gorm.Model
	Question string `gorm:"type:text"`
	Email    string
	Phone    string
	Status   int
}

// GetLanguage get language
func GetLanguage(src string) (lang string) {
	s := strings.ToLower(src)
	switch s {
	case "en", "en-us":
		lang = "en"
	case "zh-cn":
		lang = "zh-cn"
	default:
		lang = "en"
	}
	if strings.HasPrefix(src, "ar") {
		lang = "ar"
	}
	return
}
