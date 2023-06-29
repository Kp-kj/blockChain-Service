package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserCryptominerModel = (*customUserCryptominerModel)(nil)

type (
	// UserCryptominerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCryptominerModel.
	UserCryptominerModel interface {
		userCryptominerModel
	}

	customUserCryptominerModel struct {
		*defaultUserCryptominerModel
	}
)

// NewUserCryptominerModel returns a model for the database table.
func NewUserCryptominerModel(conn sqlx.SqlConn) UserCryptominerModel {
	return &customUserCryptominerModel{
		defaultUserCryptominerModel: newUserCryptominerModel(conn),
	}
}
