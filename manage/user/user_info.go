package user

import (
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// TokenCheck check token
func TokenCheck(token string) (existed bool, recorder *models.UserLoginRecorder) {
	db := common.DB
	recorder = &models.UserLoginRecorder{}
	err := db.Preload("User").Where(&models.UserLoginRecorder{Token: token}).First(recorder).Error
	if err != nil {
		return false, nil
	}
	return true, recorder
}
