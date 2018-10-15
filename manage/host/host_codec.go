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
	Level    int             `json:"level"`
	Package  int             `json:"package"`
}

// VPNHostConfig vpn host config
type VPNHostConfig struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
	Icon string `json:"icon"`
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
	if res.Start < 0 {
		res.Start = 0
	}
	res.End = conf.End.Unix()
	if res.End < 0 {
		res.End = 0
	}
	res.Level = conf.User.Level
	res.Package = conf.User.PackageType

	hosts := make([]VPNHostConfig, len(conf.Hosts))
	for i, item := range conf.Hosts {
		hosts[i] = VPNHostConfig{
			ID:   item.ID,
			Name: item.Name,
			IP:   item.IP,
			Icon: item.Icon,
		}
	}
	res.Hosts = hosts
	return
}
