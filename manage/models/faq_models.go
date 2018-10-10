package models

import "github.com/jinzhu/gorm"

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
