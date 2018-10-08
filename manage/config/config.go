package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// rune mode
const (
	DEV  = "dev"
	PROD = "prod"
)

var (
	// AppConfig is the project config
	AppConfig Config
)

// Config is config information
type Config struct {
	Mode       string
	Auth       AuthConfig
	MYSQL      MYSQLConfig
	SMTP       SMTPConfig
	AppleStore AppleStoreConfig `yaml:"appleStore"`
}

// AuthConfig is auth to api config
type AuthConfig struct {
	User     string
	Password string
}

// MYSQLConfig is mysql client information
type MYSQLConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	DB          string
	MaxIdle     int
	MaxConnect  int
	MaxLifeTime int
}

// BuildClientURI is make a mysql client uri for go
func (c MYSQLConfig) BuildClientURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
		"Asia%2FShanghai",
	)
}

// SMTPConfig is smpt server config
type SMTPConfig struct {
	User     string
	Password string
	Host     string
	Port     int
}

// AppleStoreConfig apple store config
type AppleStoreConfig struct {
	BundleID string `yaml:"bundleId"`
}

func init() {

	var (
		mode string
	)
	flag.StringVar(&mode, "mode", DEV, "run mode")

	var (
		authUser     string
		authPassword string
	)

	flag.StringVar(&authUser, "auth.user", "user", "auth user for manage")
	flag.StringVar(&authPassword, "auth.password", "pwd", "auth password for manage")

	var (
		mysqlHost       string
		mysqlPort       int
		mysqlUser       string
		mysqlPassword   string
		mysqlDB         string
		mysqlMaxIdle    int
		mysqlMaxConnect int
		mysqlLifeTime   int
	)

	flag.StringVar(&mysqlHost, "mysql.host", "127.0.0.1", "mysql host.")
	flag.IntVar(&mysqlPort, "mysql.port", 3306, "mysql port")
	flag.StringVar(&mysqlUser, "mysql.user", "demo", "mysql user")
	flag.StringVar(&mysqlPassword, "mysql.password", "pwd", "mysql password")
	flag.StringVar(&mysqlDB, "mysql.db", "test", "mysql database name")
	flag.IntVar(&mysqlMaxIdle, "mysql.maxidle", 4, "mysql max idle")
	flag.IntVar(&mysqlMaxConnect, "mysql.maxconnect", 200, "mysql max connect")
	flag.IntVar(&mysqlLifeTime, "mysql.lifetime", 0, "mysql life time")

	var (
		smtpUser     string
		smtpPassword string
		smtpHost     string
		smtpPort     int
	)
	flag.StringVar(&smtpUser, "smtp.user", "your@example.com", "email user.")
	flag.StringVar(&smtpPassword, "smtp.password", "pwd", "email password.")
	flag.StringVar(&smtpHost, "smtp.host", "stmp.163.com", "smtp server host.")
	flag.IntVar(&smtpPort, "smtp.port", 25, "smtp server port")

	var (
		appleBundleID string
	)
	flag.StringVar(&appleBundleID, "apple.bundle", "", "apple store bundle id.")

	var (
		configFile string
	)

	flag.StringVar(&configFile, "config", "", "config file.")

	flag.Parse()
	// check

	if configFile != "" {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatalln(err)
		}
		err = yaml.Unmarshal(data, &AppConfig)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%+v", AppConfig)
		return
	}
	AppConfig = Config{
		Mode: mode,
		Auth: AuthConfig{
			User:     authUser,
			Password: authPassword,
		},
		MYSQL: MYSQLConfig{
			Host:        mysqlHost,
			Port:        mysqlPort,
			User:        mysqlUser,
			Password:    mysqlPassword,
			DB:          mysqlDB,
			MaxIdle:     mysqlMaxIdle,
			MaxConnect:  mysqlMaxConnect,
			MaxLifeTime: mysqlLifeTime,
		},
		SMTP: SMTPConfig{
			User:     smtpUser,
			Password: smtpPassword,
			Host:     smtpHost,
			Port:     smtpPort,
		},
		AppleStore: AppleStoreConfig{
			BundleID: appleBundleID,
		},
	}
	fmt.Printf("%+v", AppConfig)
}
