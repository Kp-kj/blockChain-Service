package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BargainRuleModel = (*customBargainRuleModel)(nil)

type (
	// BargainRuleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBargainRuleModel.
	BargainRuleModel interface {
		bargainRuleModel
	}

	customBargainRuleModel struct {
		*defaultBargainRuleModel
	}
)

// NewBargainRuleModel returns a model for the database table.
func NewBargainRuleModel(conn sqlx.SqlConn) BargainRuleModel {
	return &customBargainRuleModel{
		defaultBargainRuleModel: newBargainRuleModel(conn),
	}
}
