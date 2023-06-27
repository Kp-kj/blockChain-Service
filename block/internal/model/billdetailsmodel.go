package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillDetailsModel = (*customBillDetailsModel)(nil)

type (
	// BillDetailsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillDetailsModel.
	BillDetailsModel interface {
		billDetailsModel
	}

	customBillDetailsModel struct {
		*defaultBillDetailsModel
	}
)

// NewBillDetailsModel returns a model for the database table.
func NewBillDetailsModel(conn sqlx.SqlConn) BillDetailsModel {
	return &customBillDetailsModel{
		defaultBillDetailsModel: newBillDetailsModel(conn),
	}
}
