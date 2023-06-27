package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BargainAmountModel = (*customBargainAmountModel)(nil)

type (
	// BargainAmountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBargainAmountModel.
	BargainAmountModel interface {
		bargainAmountModel
	}

	customBargainAmountModel struct {
		*defaultBargainAmountModel
	}
)

// NewBargainAmountModel returns a model for the database table.
func NewBargainAmountModel(conn sqlx.SqlConn) BargainAmountModel {
	return &customBargainAmountModel{
		defaultBargainAmountModel: newBargainAmountModel(conn),
	}
}
