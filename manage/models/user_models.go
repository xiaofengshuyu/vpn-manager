package models

import (
	"time"
)

// CommonUser user
type CommonUser struct {
	ID         int64
	UserName   string
	RealName   string
	Password   string
	Email      string
	Phone      string
	Status     int
	CreateTime string
	UpdateTime string
}

// UserVPNConfig user vpn config
type UserVPNConfig struct {
	ID        int64
	StartDate time.Time
	EndDate   time.Time
}
