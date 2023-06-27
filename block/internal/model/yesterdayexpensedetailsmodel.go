package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ YesterdayExpenseDetailsModel = (*customYesterdayExpenseDetailsModel)(nil)

type (
	// YesterdayExpenseDetailsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYesterdayExpenseDetailsModel.
	YesterdayExpenseDetailsModel interface {
		yesterdayExpenseDetailsModel
	}

	customYesterdayExpenseDetailsModel struct {
		*defaultYesterdayExpenseDetailsModel
	}
)

// NewYesterdayExpenseDetailsModel returns a model for the database table.
func NewYesterdayExpenseDetailsModel(conn sqlx.SqlConn) YesterdayExpenseDetailsModel {
	return &customYesterdayExpenseDetailsModel{
		defaultYesterdayExpenseDetailsModel: newYesterdayExpenseDetailsModel(conn),
	}
}
