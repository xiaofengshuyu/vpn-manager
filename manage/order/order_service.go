package order

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/host"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// Service is a interface for order service
type Service interface {
	GetProduct(ctx context.Context) (products []*models.Product, err error)
	CommitAnOrder(ctx context.Context, data string) (err error)
}

// BaseOrderService is a order service
type BaseOrderService struct{}

// GetProduct get all product
func (s *BaseOrderService) GetProduct(ctx context.Context) (products []*models.Product, err error) {
	err = common.DB.Model(&models.Product{}).Find(&products).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	return
}

// CommitAnOrder is create a order
func (s *BaseOrderService) CommitAnOrder(ctx context.Context, data string) (err error) {

	// check order from apple server
	userOrder, err := GetOrderFromApple(data)
	if err != nil {
		return
	}

	db := common.DB.Begin()
	defer func(err *error) {
		if *err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}(&err)

	// check order info
	var orders []*models.Order
	err = db.Where(&models.Order{OrderNumber: userOrder.OrderNumber}).Find(&orders).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	if len(orders) > 0 {
		err = common.NewInsertRepeatError("the order is existed")
		return
	}
	user := &models.CommonUser{}
	err = db.Where(&models.CommonUser{Email: userOrder.User.Email}).First(user).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}

	// write order info to db
	userOrder.UserID = user.ID
	userOrder.User = *user
	err = db.Create(&userOrder).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}

	// TODO
	// update user vpn config
	// get user current status
	var (
		now   = time.Now()
		month = userOrder.Product.Duration * userOrder.Quantity
	)
	var currentConfig models.UserVPNConfig
	errs := db.Where(&models.UserVPNConfig{UserID: user.ID}).First(&currentConfig).Error
	if errs != nil {
		if errs == gorm.ErrRecordNotFound {
			// create an new
			userConf := &models.UserVPNConfig{
				UserID: user.ID,
				Start:  now,
				End:    now.AddDate(0, month, 0),
			}
			errs = db.Create(userConf).Error
			if errs != nil {
				err = errs
				return
			}
		} else {
			err = errs
			return
		}
	} else {
		if currentConfig.End.After(now) {
			currentConfig.End = currentConfig.End.AddDate(0, month, 0)
		} else {
			currentConfig.Start = now
			currentConfig.End = now.AddDate(0, month, 0)
		}
		// update
		errs := db.Model(&currentConfig).Update("start", currentConfig.Start, "end", currentConfig.End).Error
		if errs != nil {
			err = errs
			return
		}
	}

	// write vpn config
	go func() {
		errs := host.AppendConfig()
		if errs != nil {
			common.Logger.Error(err)
		}
	}()
	return
}
