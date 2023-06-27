package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PurchaseRecordsModel = (*customPurchaseRecordsModel)(nil)

type (
	// PurchaseRecordsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPurchaseRecordsModel.
	PurchaseRecordsModel interface {
		purchaseRecordsModel
	}

	customPurchaseRecordsModel struct {
		*defaultPurchaseRecordsModel
	}
)

// NewPurchaseRecordsModel returns a model for the database table.
func NewPurchaseRecordsModel(conn sqlx.SqlConn) PurchaseRecordsModel {
	return &customPurchaseRecordsModel{
		defaultPurchaseRecordsModel: newPurchaseRecordsModel(conn),
	}
}
