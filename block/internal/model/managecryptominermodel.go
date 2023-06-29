package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ManageCryptominerModel = (*customManageCryptominerModel)(nil)

type (
	// ManageCryptominerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageCryptominerModel.
	ManageCryptominerModel interface {
		manageCryptominerModel
	}

	customManageCryptominerModel struct {
		*defaultManageCryptominerModel
	}
)

// NewManageCryptominerModel returns a model for the database table.
func NewManageCryptominerModel(conn sqlx.SqlConn) ManageCryptominerModel {
	return &customManageCryptominerModel{
		defaultManageCryptominerModel: newManageCryptominerModel(conn),
	}
}
