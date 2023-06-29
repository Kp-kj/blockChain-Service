package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ManageGoodsModel = (*customManageGoodsModel)(nil)

type (
	// ManageGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageGoodsModel.
	ManageGoodsModel interface {
		manageGoodsModel
	}

	customManageGoodsModel struct {
		*defaultManageGoodsModel
	}
)

// NewManageGoodsModel returns a model for the database table.
func NewManageGoodsModel(conn sqlx.SqlConn) ManageGoodsModel {
	return &customManageGoodsModel{
		defaultManageGoodsModel: newManageGoodsModel(conn),
	}
}
