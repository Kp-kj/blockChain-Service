// Code generated by goctl. DO NOT EDIT.

package myssql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	yesterdayBillFieldNames          = builder.RawFieldNames(&YesterdayBill{})
	yesterdayBillRows                = strings.Join(yesterdayBillFieldNames, ",")
	yesterdayBillRowsExpectAutoSet   = strings.Join(stringx.Remove(yesterdayBillFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	yesterdayBillRowsWithPlaceHolder = strings.Join(stringx.Remove(yesterdayBillFieldNames, "`walletId`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	yesterdayBillModel interface {
		Insert(ctx context.Context, data *YesterdayBill) (sql.Result, error)
		FindOne(ctx context.Context, walletId int64) (*YesterdayBill, error)
		Update(ctx context.Context, data *YesterdayBill) error
		Delete(ctx context.Context, walletId int64) error
	}

	defaultYesterdayBillModel struct {
		conn  sqlx.SqlConn
		table string
	}

	YesterdayBill struct {
		WalletId                 int64   `db:"walletId"`
		YesterdayRewardQuantity  float64 `db:"yesterdayRewardQuantity"`
		YesterdayExpenseQuantity float64 `db:"yesterdayExpenseQuantity"`
	}
)

func newYesterdayBillModel(conn sqlx.SqlConn) *defaultYesterdayBillModel {
	return &defaultYesterdayBillModel{
		conn:  conn,
		table: "`yesterday_bill`",
	}
}

func (m *defaultYesterdayBillModel) withSession(session sqlx.Session) *defaultYesterdayBillModel {
	return &defaultYesterdayBillModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`yesterday_bill`",
	}
}

func (m *defaultYesterdayBillModel) Delete(ctx context.Context, walletId int64) error {
	query := fmt.Sprintf("delete from %s where `walletId` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, walletId)
	return err
}

func (m *defaultYesterdayBillModel) FindOne(ctx context.Context, walletId int64) (*YesterdayBill, error) {
	query := fmt.Sprintf("select %s from %s where `walletId` = ? limit 1", yesterdayBillRows, m.table)
	var resp YesterdayBill
	err := m.conn.QueryRowCtx(ctx, &resp, query, walletId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYesterdayBillModel) Insert(ctx context.Context, data *YesterdayBill) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, yesterdayBillRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.YesterdayRewardQuantity, data.YesterdayExpenseQuantity)
	return ret, err
}

func (m *defaultYesterdayBillModel) Update(ctx context.Context, data *YesterdayBill) error {
	query := fmt.Sprintf("update %s set %s where `walletId` = ?", m.table, yesterdayBillRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.YesterdayRewardQuantity, data.YesterdayExpenseQuantity, data.WalletId)
	return err
}

func (m *defaultYesterdayBillModel) tableName() string {
	return m.table
}