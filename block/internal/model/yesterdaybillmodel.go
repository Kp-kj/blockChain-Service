package myssql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ YesterdayBillModel = (*customYesterdayBillModel)(nil)

type (
	// YesterdayBillModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYesterdayBillModel.
	YesterdayBillModel interface {
		yesterdayBillModel
	}

	customYesterdayBillModel struct {
		*defaultYesterdayBillModel
	}
)

// NewYesterdayBillModel returns a model for the database table.
func NewYesterdayBillModel(conn sqlx.SqlConn) YesterdayBillModel {
	return &customYesterdayBillModel{
		defaultYesterdayBillModel: newYesterdayBillModel(conn),
	}
}
