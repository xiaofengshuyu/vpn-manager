package config

import (
	"flag"
)

var (
	// AppConfig is the project config
	AppConfig Config
)

// Config is config information
type Config struct {
	AuthUser     string
	AuthPassword string
}

func init() {
	var (
		authUser     string
		authPassword string
	)

	flag.StringVar(&authUser, "auth-user", "user", "auth user for manage")
	flag.StringVar(&authPassword, "auth-password", "pwd", "auth password for manage")
	flag.Parse()
	// check

	AppConfig = Config{
		AuthUser:     authUser,
		AuthPassword: authPassword,
	}
}
