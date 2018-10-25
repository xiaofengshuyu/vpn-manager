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
	// Duration unit is month
	Duration int
}

// Order is user's order
type Order struct {
	gorm.Model
	UserID uint
	User   CommonUser
	// OrderNumber is OrderData's md5
	OrderNumber string `gorm:"unique;not null"`
	OrderData   string `gorm:"type:text"`
	OrderTime   string
	AddMonth    int
	Quantity    int
	Product     Product
	ProductID   uint
}
