package host

import (
	"context"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// ConfigService vpn config service interface
type ConfigService interface {
	GetVPNConfig(ctx context.Context, user *models.CommonUser) (config *models.UserVPNConfig, err error)
}

// BaseConfigService is an implement for config service
type BaseConfigService struct {
}

// GetVPNConfig return vpn config
func (s *BaseConfigService) GetVPNConfig(ctx context.Context, user *models.CommonUser) (config *models.UserVPNConfig, err error) {
	// get user info
	db := common.DB
	config = &models.UserVPNConfig{}
	err = db.Preload("Hosts").Where(&models.UserVPNConfig{UserID: user.ID}).First(config).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}

// NewConfigService create a new service
func NewConfigService() ConfigService {
	return &BaseConfigService{}
}
