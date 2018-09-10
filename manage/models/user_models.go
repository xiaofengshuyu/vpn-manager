package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

// user status
const (
	UserStatusDisable = iota
	UserStatusEnable
	UserStatusRegister
)

// CommonUser user
type CommonUser struct {
	gorm.Model
	UserName    string `validate:"required"`
	RealName    string
	Password    string `validate:"required"`
	Email       string `gorm:"unique;not null" validate:"required,email"`
	Phone       string
	Status      int `gorm:"default:1"`
	VertifyCode string
}

// TableName is mysql table name
func (CommonUser) TableName() string {
	return "t_vpn_common_user"
}

// Valid is check user
func (u *CommonUser) Valid(validate *validator.Validate) error {
	return validate.Struct(u)
}

// UserVPNConfig user vpn config
type UserVPNConfig struct {
	gorm.Model
	StartDate time.Time
	EndDate   time.Time
}
