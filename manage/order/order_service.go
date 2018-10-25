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
	user, ok := ctx.Value(common.UserInfoKey).(models.CommonUser)
	if !ok {
		err = common.NewNotLoginError("cannot read user info from context")
		return
	}

	// check order from apple server
	userOrders, err := GetOrderFromApple(data)
	if err != nil {
		return
	}

	if len(userOrders) == 0 {
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

	// get current order
	var (
		needAddOrders = make([]*models.Order, 0)
		savedOrders   = make([]*models.Order, 0)
	)
	// check order info
	err = db.Where(&models.Order{UserID: user.ID}).Find(&savedOrders).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	// compare
	for _, userOrder := range userOrders {
		find := false
		for _, old := range savedOrders {
			if userOrder.OrderNumber == old.OrderNumber {
				find = true
			}
		}
		if !find {
			needAddOrders = append(needAddOrders, &models.Order{
				UserID:      user.ID,
				OrderNumber: userOrder.OrderNumber,
				OrderData:   userOrder.OrderData,
				OrderTime:   userOrder.OrderTime,
				Quantity:    userOrder.Quantity,
				AddMonth:    userOrder.AddMonth,
				ProductID:   userOrder.ProductID,
			})
		}
	}
	if len(needAddOrders) == 0 {
		err = common.NewInsertRepeatError("all order is existed")
		return
	}

	// TODO set package type by product
	var packageType = models.UserPackageTypeCommon
	// change user level
	if user.Level == models.UserLevelFree {
		err = db.Model(user).UpdateColumns(
			models.CommonUser{
				Level:       models.UserLevelCommon,
				PackageType: packageType,
			},
		).Error
		if err != nil {
			err = common.NewDBAccessError(err)
			return
		}
	}
	var month int
	for _, item := range needAddOrders {
		month += item.AddMonth
		err = db.Create(item).Error
		if err != nil {
			err = common.NewDBAccessError(err)
			return
		}
	}

	// TODO
	// get user current status
	var (
		now = time.Now()
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
		errs := db.Model(&currentConfig).UpdateColumns(
			models.UserVPNConfig{
				Start: currentConfig.Start,
				End:   currentConfig.End,
			},
		).Error
		if errs != nil {
			err = errs
			return
		}
	}

	// write vpn config
	go func() {
		errs := host.WriteConfigToLocal()
		// errs := host.AppendConfig()
		if errs != nil {
			common.Logger.Error(err)
		}
	}()
	return
}
