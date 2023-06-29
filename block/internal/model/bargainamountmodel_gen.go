// Code generated by goctl. DO NOT EDIT.

package model

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
	bargainAmountFieldNames          = builder.RawFieldNames(&BargainAmount{})
	bargainAmountRows                = strings.Join(bargainAmountFieldNames, ",")
	bargainAmountRowsExpectAutoSet   = strings.Join(stringx.Remove(bargainAmountFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	bargainAmountRowsWithPlaceHolder = strings.Join(stringx.Remove(bargainAmountFieldNames, "`bargain_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	bargainAmountModel interface {
		Insert(ctx context.Context, data *BargainAmount) (sql.Result, error)
		FindOne(ctx context.Context, bargainId int64) (*BargainAmount, error)
		Update(ctx context.Context, data *BargainAmount) error
		Delete(ctx context.Context, bargainId int64) error
	}

	defaultBargainAmountModel struct {
		conn  sqlx.SqlConn
		table string
	}

	BargainAmount struct {
		BargainId              int64   `db:"bargain_id"`
		FirstBargainPercentage float64 `db:"first_bargain_percentage"`
		BargainMinPrice        float64 `db:"bargain_min_price"`
		BargainPrice           float64 `db:"bargain_price"`
	}
)

func newBargainAmountModel(conn sqlx.SqlConn) *defaultBargainAmountModel {
	return &defaultBargainAmountModel{
		conn:  conn,
		table: "`bargain_amount`",
	}
}

func (m *defaultBargainAmountModel) Delete(ctx context.Context, bargainId int64) error {
	query := fmt.Sprintf("delete from %s where `bargain_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, bargainId)
	return err
}

func (m *defaultBargainAmountModel) FindOne(ctx context.Context, bargainId int64) (*BargainAmount, error) {
	query := fmt.Sprintf("select %s from %s where `bargain_id` = ? limit 1", bargainAmountRows, m.table)
	var resp BargainAmount
	err := m.conn.QueryRowCtx(ctx, &resp, query, bargainId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBargainAmountModel) Insert(ctx context.Context, data *BargainAmount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, bargainAmountRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.BargainId, data.FirstBargainPercentage, data.BargainMinPrice, data.BargainPrice)
	return ret, err
}

func (m *defaultBargainAmountModel) Update(ctx context.Context, data *BargainAmount) error {
	query := fmt.Sprintf("update %s set %s where `bargain_id` = ?", m.table, bargainAmountRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.FirstBargainPercentage, data.BargainMinPrice, data.BargainPrice, data.BargainId)
	return err
}

func (m *defaultBargainAmountModel) tableName() string {
	return m.table
}
