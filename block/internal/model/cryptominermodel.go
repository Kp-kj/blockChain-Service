package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CryptominerModel = (*customCryptominerModel)(nil)

type (
	// CryptominerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCryptominerModel.
	CryptominerModel interface {
		cryptominerModel
	}

	customCryptominerModel struct {
		*defaultCryptominerModel
	}
)

// NewCryptominerModel returns a model for the database table.
func NewCryptominerModel(conn sqlx.SqlConn) CryptominerModel {
	return &customCryptominerModel{
		defaultCryptominerModel: newCryptominerModel(conn),
	}
}
