package myssql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MallsModel = (*customMallsModel)(nil)

type (
	// MallsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMallsModel.
	MallsModel interface {
		mallsModel
	}

	customMallsModel struct {
		*defaultMallsModel
	}
)

// NewMallsModel returns a model for the database table.
func NewMallsModel(conn sqlx.SqlConn) MallsModel {
	return &customMallsModel{
		defaultMallsModel: newMallsModel(conn),
	}
}
