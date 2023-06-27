package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ YesterdayRewardDetailsModel = (*customYesterdayRewardDetailsModel)(nil)

type (
	// YesterdayRewardDetailsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYesterdayRewardDetailsModel.
	YesterdayRewardDetailsModel interface {
		yesterdayRewardDetailsModel
	}

	customYesterdayRewardDetailsModel struct {
		*defaultYesterdayRewardDetailsModel
	}
)

// NewYesterdayRewardDetailsModel returns a model for the database table.
func NewYesterdayRewardDetailsModel(conn sqlx.SqlConn) YesterdayRewardDetailsModel {
	return &customYesterdayRewardDetailsModel{
		defaultYesterdayRewardDetailsModel: newYesterdayRewardDetailsModel(conn),
	}
}
