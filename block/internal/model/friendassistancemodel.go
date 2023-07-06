package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FriendAssistanceModel = (*customFriendAssistanceModel)(nil)

type (
	// FriendAssistanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendAssistanceModel.
	FriendAssistanceModel interface {
		friendAssistanceModel
	}

	customFriendAssistanceModel struct {
		*defaultFriendAssistanceModel
	}
)

// NewFriendAssistanceModel returns a model for the database table.
func NewFriendAssistanceModel(conn sqlx.SqlConn) FriendAssistanceModel {
	return &customFriendAssistanceModel{
		defaultFriendAssistanceModel: newFriendAssistanceModel(conn),
	}
}
