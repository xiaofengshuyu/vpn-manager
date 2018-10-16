package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
	"github.com/xiaofengshuyu/vpn-manager/manage/utils"
)

// Service is a interface for user service
type Service interface {
	// AddUser()
	// UpdateUser()
	// DeleteUser()
	GetUserOne(ctx context.Context, cond *models.CommonUser) (user *models.CommonUser, err error)
	RegisterUserWithoutCheck(ctx context.Context, user *models.CommonUser) (err error)
	// GetUserByID()
	RegisterUser(ctx context.Context, user *models.CommonUser) (err error)
	EmailResend(ctx context.Context, user *models.CommonUser) (err error)
	RegisterCheck(ctx context.Context, user *models.CommonUser) (err error)
	Login(ctx context.Context, username, password string) (recorder *models.UserLoginRecorder, err error)

	ResetPassword(ctx context.Context, user *models.CommonUser) (err error)
}

// BaseUserService is a implements for user service
type BaseUserService struct{}

// RegisterUser is register user
func (s *BaseUserService) RegisterUser(ctx context.Context, user *models.CommonUser) (err error) {
	db := common.DB.Begin()
	defer func(err *error) {
		if *err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}(&err)
	var users []*models.CommonUser

	// check is existed
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
	user.VertifyCodeStart = now

	err = db.Create(user).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	// send an email
	go func() {
		// TODO email templement
		errs := utils.SendSimpleEmail([]string{user.Email}, "Vertify Code", user.VertifyCode)
		if errs != nil {
			logger.Error(errs)
		}
	}()
	return
}

// RegisterUserWithoutCheck is register without vertify code
func (s *BaseUserService) RegisterUserWithoutCheck(ctx context.Context, user *models.CommonUser) (err error) {
	db := common.DB.Begin()
	defer func(err *error) {
		if *err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}(&err)
	var users []*models.CommonUser

	// check is existed
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
	user.Status = models.UserStatusEnable
	user.CreatedAt = now
	user.UpdatedAt = now

	// create a verification code
	user.VertifyCode = makeVertifyCode()
	user.VertifyCodeStart = now

	err = db.Create(user).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}

// EmailResend email resend, user must have email attribute
func (s *BaseUserService) EmailResend(ctx context.Context, user *models.CommonUser) (err error) {
	db := common.DB.Begin()
	defer func(err *error) {
		if *err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}(&err)
	u := &models.CommonUser{}
	err = db.Where(&models.CommonUser{Email: user.Email}).First(u).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	veritifyCode := makeVertifyCode()
	err = db.Model(u).Update(
		"veritify_code", veritifyCode,
		"veritify_code_start", time.Now(),
	).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	// send an email
	go func() {
		// TODO email templement
		errs := utils.SendSimpleEmail([]string{user.Email}, "Vertify Code", veritifyCode)
		if errs != nil {
			logger.Error(errs)
		} else {
			logger.Info(fmt.Sprintf("send email to %s succeed,vertify code is %s.", user.Email, veritifyCode))
		}
	}()
	return
}

// RegisterCheck check register info,user must have email and veritifyCode
func (s *BaseUserService) RegisterCheck(ctx context.Context, user *models.CommonUser) (err error) {
	// check user and vertify code
	db := common.DB.Begin()
	defer func(err *error) {
		if *err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}(&err)
	u := &models.CommonUser{}
	err = db.Where(&models.CommonUser{Email: user.Email}).First(u).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	if u.VertifyCode != user.VertifyCode || !u.VertifyCodeIsValid() {
		err = common.NewRequestParamsValueError(ErrVertifyCodeInvalid)
		return
	}
	err = db.Model(u).Update("status", models.UserStatusEnable).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}

// GetUserOne get a user by condition
func (s *BaseUserService) GetUserOne(ctx context.Context, cond *models.CommonUser) (user *models.CommonUser, err error) {
	db := common.DB
	user = &models.CommonUser{}
	err = db.Where(cond).First(user).Error
	if err != nil {
		err = common.NewResourcesNotFoundError(err)
	}
	return
}

// Login login
func (s *BaseUserService) Login(ctx context.Context, username, password string) (recorder *models.UserLoginRecorder, err error) {
	db := common.DB
	user := &models.CommonUser{}
	err = db.Where("email = ? or user_name = ?", username, username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = common.NewNotRegisterError()
			return
		}
		err = common.NewDBAccessError(err)
		return
	}
	if user.Password != password {
		err = common.NewRequestParamsValueError("password error")
		return
	}
	if user.Status != models.UserStatusEnable {
		err = common.NewNotRegisterError()
		return
	}
	// write login info to recorder
	now := time.Now()
	token := makeToken(user.Email)
	refreshToken := makeToken(user.Email + user.Password)
	recorder = &models.UserLoginRecorder{
		UserID:       int(user.ID),
		Token:        token,
		RefreshToken: refreshToken,
		LastLogin:    now,
		EndTime:      now.AddDate(0, 1, 0),
	}
	// check is existed
	old := &models.UserLoginRecorder{UserID: int(user.ID)}
	err = db.Where(old).First(old).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = db.Create(recorder).Error
			if err != nil {
				err = common.NewDBAccessError(err)
				return
			}
			return
		}
		err = common.NewDBAccessError(err)
		return
	}
	recorder.ID = old.ID
	err = db.Save(recorder).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}

// ResetPassword reset password
func (s *BaseUserService) ResetPassword(ctx context.Context, user *models.CommonUser) (err error) {
	db := common.DB.Begin()
	defer func(err *error) {
		if *err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}(&err)
	u := &models.CommonUser{}
	err = db.Where(&models.CommonUser{Email: user.Email}).First(u).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	if u.Status != models.UserStatusEnable {
		err = common.NewRequestParamsValueError("user is not enable")
		return
	}
	if u.VertifyCode != user.VertifyCode || !u.VertifyCodeIsValid() {
		err = common.NewRequestParamsValueError(ErrVertifyCodeInvalid)
		return
	}
	err = db.Model(u).Update("password", user.Password).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}
