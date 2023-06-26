package myssql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProfileModel = (*customProfileModel)(nil)

type (
	// ProfileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProfileModel.
	ProfileModel interface {
		profileModel
	}

	customProfileModel struct {
		*defaultProfileModel
	}
)

// NewProfileModel returns a model for the database table.
func NewProfileModel(conn sqlx.SqlConn) ProfileModel {
	return &customProfileModel{
		defaultProfileModel: newProfileModel(conn),
	}
}
