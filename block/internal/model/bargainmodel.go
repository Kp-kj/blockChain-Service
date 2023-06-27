package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BargainModel = (*customBargainModel)(nil)

type (
	// BargainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBargainModel.
	BargainModel interface {
		bargainModel
	}

	customBargainModel struct {
		*defaultBargainModel
	}
)

// NewBargainModel returns a model for the database table.
func NewBargainModel(conn sqlx.SqlConn) BargainModel {
	return &customBargainModel{
		defaultBargainModel: newBargainModel(conn),
	}
}
