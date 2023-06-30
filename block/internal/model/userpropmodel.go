package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserPropModel = (*customUserPropModel)(nil)

type (
	// UserPropModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPropModel.
	UserPropModel interface {
		userPropModel
	}

	customUserPropModel struct {
		*defaultUserPropModel
	}
)

// NewUserPropModel returns a model for the database table.
func NewUserPropModel(conn sqlx.SqlConn) UserPropModel {
	return &customUserPropModel{
		defaultUserPropModel: newUserPropModel(conn),
	}
}
