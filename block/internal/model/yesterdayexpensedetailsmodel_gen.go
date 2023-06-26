// Code generated by goctl. DO NOT EDIT.

package myssql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	yesterdayExpenseDetailsFieldNames          = builder.RawFieldNames(&YesterdayExpenseDetails{})
	yesterdayExpenseDetailsRows                = strings.Join(yesterdayExpenseDetailsFieldNames, ",")
	yesterdayExpenseDetailsRowsExpectAutoSet   = strings.Join(stringx.Remove(yesterdayExpenseDetailsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	yesterdayExpenseDetailsRowsWithPlaceHolder = strings.Join(stringx.Remove(yesterdayExpenseDetailsFieldNames, "`walletId`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	yesterdayExpenseDetailsModel interface {
		Insert(ctx context.Context, data *YesterdayExpenseDetails) (sql.Result, error)
		FindOne(ctx context.Context, walletId int64) (*YesterdayExpenseDetails, error)
		Update(ctx context.Context, data *YesterdayExpenseDetails) error
		Delete(ctx context.Context, walletId int64) error
	}

	defaultYesterdayExpenseDetailsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	YesterdayExpenseDetails struct {
		WalletId      int64     `db:"walletId"`
		ExpenseName   string    `db:"expenseName"`
		ItemQuantity  int64     `db:"itemQuantity"`
		ExpenseAmount float64   `db:"expenseAmount"`
		ExpenseTime   time.Time `db:"expenseTime"`
	}
)

func newYesterdayExpenseDetailsModel(conn sqlx.SqlConn) *defaultYesterdayExpenseDetailsModel {
	return &defaultYesterdayExpenseDetailsModel{
		conn:  conn,
		table: "`yesterday_expense_details`",
	}
}

func (m *defaultYesterdayExpenseDetailsModel) withSession(session sqlx.Session) *defaultYesterdayExpenseDetailsModel {
	return &defaultYesterdayExpenseDetailsModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`yesterday_expense_details`",
	}
}

func (m *defaultYesterdayExpenseDetailsModel) Delete(ctx context.Context, walletId int64) error {
	query := fmt.Sprintf("delete from %s where `walletId` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, walletId)
	return err
}

func (m *defaultYesterdayExpenseDetailsModel) FindOne(ctx context.Context, walletId int64) (*YesterdayExpenseDetails, error) {
	query := fmt.Sprintf("select %s from %s where `walletId` = ? limit 1", yesterdayExpenseDetailsRows, m.table)
	var resp YesterdayExpenseDetails
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

func (m *defaultYesterdayExpenseDetailsModel) Insert(ctx context.Context, data *YesterdayExpenseDetails) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, yesterdayExpenseDetailsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.WalletId, data.ExpenseName, data.ItemQuantity, data.ExpenseAmount, data.ExpenseTime)
	return ret, err
}

func (m *defaultYesterdayExpenseDetailsModel) Update(ctx context.Context, data *YesterdayExpenseDetails) error {
	query := fmt.Sprintf("update %s set %s where `walletId` = ?", m.table, yesterdayExpenseDetailsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ExpenseName, data.ItemQuantity, data.ExpenseAmount, data.ExpenseTime, data.WalletId)
	return err
}

func (m *defaultYesterdayExpenseDetailsModel) tableName() string {
	return m.table
}
