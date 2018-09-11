package user

import (
	"context"
	"time"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// Service is a interface for user service
type Service interface {
	// AddUser()
	// UpdateUser()
	// DeleteUser()
	// GetUser()
	// GetUserByID()
	RegisterUser(ctx context.Context, user *models.CommonUser) (err error)
	RegisterCheck()
}

// BaseUserService is a implements for user service
type BaseUserService struct{}

// RegisterUser is register user
func (s *BaseUserService) RegisterUser(ctx context.Context, user *models.CommonUser) (err error) {
	// check is existed
	db := common.DB.Begin()
	defer func(err *error) {
		if *err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}(&err)
	var users []*models.CommonUser

	err = db.Where(&models.CommonUser{Email: user.Email}).Find(&users).Error
	if err != nil {
		common.NewDBAccessError(err)
		return
	}
	if len(users) > 0 {
		err = common.NewInsertRepeatError("email is existed")
		return
	}

	user.UserName = user.Email
	now := time.Now()
	user.Status = models.UserStatusRegister
	user.CreatedAt = now
	user.UpdatedAt = now

	// create a verification code
	user.VertifyCode = makeVertifyCode()

	err = db.Create(user).Error
	if err != nil {
		err = common.NewDBAccessError(err)
	}
	return
}

// RegisterCheck check register info
func (s *BaseUserService) RegisterCheck() {
	return
}
