package order

import (
	"context"

	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/host"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

// Service is a interface for order service
type Service interface {
	GetProduct(ctx context.Context) (products []*models.Product, err error)
	CommitAnOrder(ctx context.Context, userOrder *models.Order) (err error)
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
func (s *BaseOrderService) CommitAnOrder(ctx context.Context, userOrder *models.Order) (err error) {
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
	product := &models.Product{}
	err = db.Where(&models.Product{Code: userOrder.Product.Code}).First(product).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}
	user := &models.CommonUser{}
	err = db.Where(&models.CommonUser{Email: userOrder.User.Email}).First(user).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}

	// check order from apple server
	// TODO
	GetOrderFromApple(userOrder.OrderNumber)

	// write order info to db
	userOrder.ProductID = int(product.ID)
	userOrder.Product = *product
	userOrder.UserID = int(user.ID)
	userOrder.User = *user
	err = db.Create(userOrder).Error
	if err != nil {
		err = common.NewDBAccessError(err)
		return
	}

	// write vpn config
	go func() {
		errs := host.AppendConfig(userOrder)
		if errs != nil {
			common.Logger.Error(err)
		}
	}()
	return
}
