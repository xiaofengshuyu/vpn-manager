package host

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"time"

	"github.com/pkg/sftp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
	"golang.org/x/crypto/ssh"
)

// AppendConfig append vpn config for user
func AppendConfig() (err error) {
	// TODO
	// send diffent config to server
	var hosts []*models.Host
	db := common.DB
	// get all hosts
	err = db.Where("status = ? and type = ?", models.HostEnable, models.HostTypeCommon).Find(&hosts).Error
	if err != nil {
		return
	}
	// build user secrets file data
	l2TPData, xAuthData, err := makeUserAuthData()
	if err != nil {
		return
	}
	for _, host := range hosts {
		// get client
		client, errs := getClientWithPassword(host.IP, host.Port, host.Username, host.Password)
		if errs != nil {
			common.Logger.Error(errs)
			continue
		}
		// copy file
		// copy L2TP file
		dst, errs := client.Create(host.L2TPFile)
		if errs != nil {
			common.Logger.Error(errs)
			client.Close()
			continue
		}
		_, errs = dst.Write(l2TPData)
		if errs != nil {
			common.Logger.Error(errs)
			client.Close()
			continue
		}
		// copy XAuth file
		dst, errs = client.Create(host.XAuthFile)
		if errs != nil {
			common.Logger.Error(errs)
			client.Close()
			continue
		}
		_, errs = dst.Write(xAuthData)
		if errs != nil {
			common.Logger.Error(errs)
			client.Close()
			continue
		}
		client.Close()
	}
	return
}

func makeUserAuthData() (l2TPData []byte, xAuthData []byte, err error) {
	// read data from db
	db := common.DB
	var configs []models.UserVPNConfig
	err = db.Preload("User").Where("end > ?", time.Now()).Find(&configs).Error
	if err != nil {
		return
	}
	l2TPBuf := &bytes.Buffer{}
	xAuthBuf := &bytes.Buffer{}
	for _, item := range configs {
		l2TPBuf.WriteString(makeL2TPSecert(item.User.Email, item.User.Password))
		xAuthBuf.WriteString(makeXAuthSecert(item.User.Email, item.User.Password))
	}
	return l2TPBuf.Bytes(), xAuthBuf.Bytes(), nil
}

func makeL2TPSecert(user, password string) string {
	return fmt.Sprintf("%s  l2tpd  %s  *\n", user, password)
}

func makeXAuthSecert(user, password string) string {
	pdata, _ := exec.Command("openssl", "passwd", "-1", password).Output()
	return fmt.Sprintf("%s:%s:xauth-psk\n", user, pdata)
}

func getClientWithPassword(host string, port int, user, password string) (client *sftp.Client, err error) {
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil },
	})
	if err != nil {
		return
	}
	return sftp.NewClient(sshClient)
}
