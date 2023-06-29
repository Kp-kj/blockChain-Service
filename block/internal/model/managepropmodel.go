package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ManagePropModel = (*customManagePropModel)(nil)

type (
	// ManagePropModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManagePropModel.
	ManagePropModel interface {
		managePropModel
	}

	customManagePropModel struct {
		*defaultManagePropModel
	}
)

// NewManagePropModel returns a model for the database table.
func NewManagePropModel(conn sqlx.SqlConn) ManagePropModel {
	return &customManagePropModel{
		defaultManagePropModel: newManagePropModel(conn),
	}
}
