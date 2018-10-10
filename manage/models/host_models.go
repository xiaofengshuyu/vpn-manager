package models

import (
	"github.com/jinzhu/gorm"
)

// host status
const (
	HostEnable  = 1
	HostDisable = 0
)

// host type
const (
	HostTypeCommon = iota
	HostTypeFree
	HostTypeSelf
)

// Host is machine information
type Host struct {
	gorm.Model
	Status   int
	Name     string
	Type     int
	IP       string `gorm:"column(ip)"`
	Port     int
	Location string
	// IPsecPSK IPsecPSK value
	IPsecPSK string `gorm:"column(ipsec_psk);type:text"`
	// IPsecPSKPath file path
	IPsecPSKPath string `gorm:"column(ipsec_psk_path)"`
	Username     string
	Password     string
	ConfigPath   string
	// L2tpFile remote secrets file which save
	// vpn user and password
	L2TPFile string `gorm:"column(l2tp_file)"`
	// XauthFile remote secrets file which save
	// vpn user and password
	XAuthFile string `gorm:"column(xauth_file)"`
}
