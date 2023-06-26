package myssql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ WalletsModel = (*customWalletsModel)(nil)

type (
	// WalletsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWalletsModel.
	WalletsModel interface {
		walletsModel
	}

	customWalletsModel struct {
		*defaultWalletsModel
	}
)

// NewWalletsModel returns a model for the database table.
func NewWalletsModel(conn sqlx.SqlConn) WalletsModel {
	return &customWalletsModel{
		defaultWalletsModel: newWalletsModel(conn),
	}
}
