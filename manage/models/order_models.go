package models

import "github.com/jinzhu/gorm"

// product type
const (
	_ = iota
	ProductTypeDay
	ProductTypeMonth
	ProductTypeYear
)

// Product is vpn productor
type Product struct {
	gorm.Model
	Name  string
	Code  string `gorm:"unique;not null"`
	Type  int
	Price float64
}

// Order is user's order
type Order struct {
	gorm.Model
	UserID      int
	User        CommonUser
	OrderNumber string `gorm:"unique;not null"`
	Product     Product
	ProductID   int
}
