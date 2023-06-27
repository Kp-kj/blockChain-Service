package svc

import (
	"block/internal/config"
	"block/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	ManageGoodsModel model.ManageGoodsModel // 商城
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ManageGoodsModel: model.NewManageGoodsModel(sqlx.NewMysql(c.DB.DataSource)), // 商城
	}
}
