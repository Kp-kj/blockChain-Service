package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ChainConfigModel = (*customChainConfigModel)(nil)

type (
	// ChainConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChainConfigModel.
	ChainConfigModel interface {
		chainConfigModel
	}

	customChainConfigModel struct {
		*defaultChainConfigModel
	}
)

// NewChainConfigModel returns a model for the database table.
func NewChainConfigModel(conn sqlx.SqlConn) ChainConfigModel {
	return &customChainConfigModel{
		defaultChainConfigModel: newChainConfigModel(conn),
	}
}
