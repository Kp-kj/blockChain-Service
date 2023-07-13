package svc

import (
	"block/internal/config"
	"block/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                 config.Config
	BargainModel           model.BargainModel
	ManageCryptominerModel model.ManageCryptominerModel
	ManagePropModel        model.ManagePropModel
	ManageActivityModel    model.ManageActivityModel
	CryptominerModel       model.CryptominerModel
	PropModel              model.PropModel
	PurchaseRecordsModel   model.PurchaseRecordsModel
	BargainRuleModel       model.BargainRuleModel
	FriendAssistanceModel  model.FriendAssistanceModel
	ChainConfigModel       model.ChainConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		BargainModel:           model.NewBargainModel(sqlx.NewMysql(c.DB.DataSource)),
		ManageCryptominerModel: model.NewManageCryptominerModel(sqlx.NewMysql(c.DB.DataSource)),
		ManagePropModel:        model.NewManagePropModel(sqlx.NewMysql(c.DB.DataSource)),
		ManageActivityModel:    model.NewManageActivityModel(sqlx.NewMysql(c.DB.DataSource)),
		CryptominerModel:       model.NewCryptominerModel(sqlx.NewMysql(c.DB.DataSource)),
		PropModel:              model.NewPropModel(sqlx.NewMysql(c.DB.DataSource)),
		PurchaseRecordsModel:   model.NewPurchaseRecordsModel(sqlx.NewMysql(c.DB.DataSource)),
		BargainRuleModel:       model.NewBargainRuleModel(sqlx.NewMysql(c.DB.DataSource)),
		FriendAssistanceModel:  model.NewFriendAssistanceModel(sqlx.NewMysql(c.DB.DataSource)),
		ChainConfigModel:       model.NewChainConfigModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
