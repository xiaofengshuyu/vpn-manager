package host

import (
	"time"

	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// VPNConfigResponse vpn config
type VPNConfigResponse struct {
	Enable   bool            `json:"enable"`
	UserName string          `json:"userName"`
	Password string          `json:"password"`
	Hosts    []VPNHostConfig `json:"hosts"`
	Start    int64           `json:"start"`
	End      int64           `json:"end"`
}

// VPNHostConfig vpn host config
type VPNHostConfig struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func vpnConfigResponseEncode(conf *models.UserVPNConfig) (res VPNConfigResponse) {
	if conf == nil {
		res.Enable = false
		return
	}
	if conf.End.After(time.Now()) {
		res.Enable = true
	}
	res.UserName = conf.User.Email
	res.Password = conf.User.Password
	res.Start = conf.Start.Unix()
	res.End = conf.End.Unix()
	hosts := make([]VPNHostConfig, len(conf.Hosts))
	for i, item := range conf.Hosts {
		hosts[i] = VPNHostConfig{
			Name: item.Name,
			IP:   item.IP,
		}
	}
	res.Hosts = hosts
	return
}
