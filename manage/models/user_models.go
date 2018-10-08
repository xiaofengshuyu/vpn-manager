package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// user status
const (
	UserStatusDisable = iota
	UserStatusRegister
	UserStatusEnable
)

// CommonUser user
type CommonUser struct {
	gorm.Model
	UserName         string `validate:"required"`
	RealName         string
	Password         string `validate:"required"`
	Email            string `gorm:"unique;not null" validate:"required,email"`
	Phone            string
	Status           int `gorm:"default:0"`
	VertifyCode      string
	VertifyCodeStart time.Time
}

// TableName is mysql table name
func (CommonUser) TableName() string {
	return "t_vpn_common_user"
}

// VertifyCodeIsValid vertify is valid, is not timeout
func (u *CommonUser) VertifyCodeIsValid() bool {
	return u.VertifyCodeStart.Add(10 * time.Minute).After(time.Now())
}

// Validate is check user
func (u *CommonUser) Validate() string {
	return TranslateAll(*u)
}

// UserLoginRecorder is login info
type UserLoginRecorder struct {
	ID           uint `gorm:"primary_key"`
	UserID       int
	User         CommonUser
	Token        string
	RefreshToken string
	EndTime      time.Time
	LastLogin    time.Time
}

// TableName is mysql table name
func (UserLoginRecorder) TableName() string {
	return "t_vpn_login_recoreder"
}

// UserVPNConfig user's vpn time config
type UserVPNConfig struct {
	gorm.Model
	UserID uint
	User   CommonUser
	Hosts  []Host
	Start  time.Time
	End    time.Time
}
