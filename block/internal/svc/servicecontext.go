package svc

import (
	"block/internal/config"
	"block/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                       config.Config
	BargainAmountModel           model.BargainAmountModel
	BargainModel                 model.BargainModel
	BillDetailsModel             model.BillDetailsModel
	ManageCryptominerModel       model.ManageCryptominerModel
	ManagePropModel              model.ManagePropModel
	ManageActivityModel          model.ManageActivityModel
	CryptominerModel             model.CryptominerModel
	PropModel                    model.PropModel
	UserCryptominerModel         model.UserCryptominerModel
	UserPropModel                model.UserPropModel
	PurchaseRecordsModel         model.PurchaseRecordsModel
	WalletsModel                 model.WalletsModel
	YesterdayBillModel           model.YesterdayBillModel
	YesterdayExpenseDetailsModel model.YesterdayExpenseDetailsModel
	YesterdayRewardDetailsModel  model.YesterdayRewardDetailsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                       c,
		BargainAmountModel:           model.NewBargainAmountModel(sqlx.NewMysql(c.DB.DataSource)),
		BargainModel:                 model.NewBargainModel(sqlx.NewMysql(c.DB.DataSource)),
		BillDetailsModel:             model.NewBillDetailsModel(sqlx.NewMysql(c.DB.DataSource)),
		ManageCryptominerModel:       model.NewManageCryptominerModel(sqlx.NewMysql(c.DB.DataSource)),
		ManagePropModel:              model.NewManagePropModel(sqlx.NewMysql(c.DB.DataSource)),
		ManageActivityModel:          model.NewManageActivityModel(sqlx.NewMysql(c.DB.DataSource)),
		CryptominerModel:             model.NewCryptominerModel(sqlx.NewMysql(c.DB.DataSource)),
		PropModel:                    model.NewPropModel(sqlx.NewMysql(c.DB.DataSource)),
		UserCryptominerModel:         model.NewUserCryptominerModel(sqlx.NewMysql(c.DB.DataSource)),
		UserPropModel:                model.NewUserPropModel(sqlx.NewMysql(c.DB.DataSource)),
		PurchaseRecordsModel:         model.NewPurchaseRecordsModel(sqlx.NewMysql(c.DB.DataSource)),
		WalletsModel:                 model.NewWalletsModel(sqlx.NewMysql(c.DB.DataSource)),
		YesterdayBillModel:           model.NewYesterdayBillModel(sqlx.NewMysql(c.DB.DataSource)),
		YesterdayExpenseDetailsModel: model.NewYesterdayExpenseDetailsModel(sqlx.NewMysql(c.DB.DataSource)),
		YesterdayRewardDetailsModel:  model.NewYesterdayRewardDetailsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
