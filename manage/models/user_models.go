package models

import (
	"time"

	"github.com/jinzhu/gorm"
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

// Validate is check user
func (u *CommonUser) Validate() string {
	return TranslateAll(*u)
}

// UserVPNConfig user vpn config
type UserVPNConfig struct {
	gorm.Model
	StartDate time.Time
	EndDate   time.Time
}
