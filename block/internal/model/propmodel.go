package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PropModel = (*customPropModel)(nil)

type (
	// PropModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPropModel.
	PropModel interface {
		propModel
	}

	customPropModel struct {
		*defaultPropModel
	}
)

// NewPropModel returns a model for the database table.
func NewPropModel(conn sqlx.SqlConn) PropModel {
	return &customPropModel{
		defaultPropModel: newPropModel(conn),
	}
}
