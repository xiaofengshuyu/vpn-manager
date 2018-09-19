package models

import (
	"github.com/jinzhu/gorm"
)

// Host is machine information
type Host struct {
	gorm.Model
	IP           string `gorm:"column(ip)"`
	Location     string
	IPsecPSK     string `gorm:"column(ipsec_psk);type:text"`
	IPsecPSKPath string `gorm:"column(ipsec_psk_path)"`
	Username     string
	Password     string
	ConfigPath   string
}
