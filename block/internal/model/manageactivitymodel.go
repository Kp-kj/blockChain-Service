package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ManageActivityModel = (*customManageActivityModel)(nil)

type (
	// ManageActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageActivityModel.
	ManageActivityModel interface {
		manageActivityModel
	}

	customManageActivityModel struct {
		*defaultManageActivityModel
	}
)

// NewManageActivityModel returns a model for the database table.
func NewManageActivityModel(conn sqlx.SqlConn) ManageActivityModel {
	return &customManageActivityModel{
		defaultManageActivityModel: newManageActivityModel(conn),
	}
}
