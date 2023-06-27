// Code generated by goctl. DO NOT EDIT.

package model

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
	manageGoodsFieldNames          = builder.RawFieldNames(&ManageGoods{})
	manageGoodsRows                = strings.Join(manageGoodsFieldNames, ",")
	manageGoodsRowsExpectAutoSet   = strings.Join(stringx.Remove(manageGoodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	manageGoodsRowsWithPlaceHolder = strings.Join(stringx.Remove(manageGoodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	manageGoodsModel interface {
		Insert(ctx context.Context, data *ManageGoods) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ManageGoods, error)
		Update(ctx context.Context, data *ManageGoods) error
		Delete(ctx context.Context, id int64) error
	}

	defaultManageGoodsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ManageGoods struct {
		Id           int64          `db:"id"`
		CreatedAt    time.Time      `db:"created_at"`
		UpdatedAt    sql.NullTime   `db:"updated_at"`
		DeletedAt    sql.NullTime   `db:"deleted_at"`
		UserId       int64          `db:"user_id"`
		GoodName     string         `db:"good_name"`
		GoodPicture  sql.NullString `db:"good_picture"`
		GoodDescribe sql.NullString `db:"good_describe"`
		GoodPrice    sql.NullInt64  `db:"good_price"`
		CurrencyType string         `db:"currency_type"`
	}
)

func newManageGoodsModel(conn sqlx.SqlConn) *defaultManageGoodsModel {
	return &defaultManageGoodsModel{
		conn:  conn,
		table: "`manage_goods`",
	}
}

func (m *defaultManageGoodsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultManageGoodsModel) FindOne(ctx context.Context, id int64) (*ManageGoods, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", manageGoodsRows, m.table)
	var resp ManageGoods
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultManageGoodsModel) Insert(ctx context.Context, data *ManageGoods) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, manageGoodsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.GoodName, data.GoodPicture, data.GoodDescribe, data.GoodPrice, data.CurrencyType)
	return ret, err
}

func (m *defaultManageGoodsModel) Update(ctx context.Context, data *ManageGoods) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, manageGoodsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.GoodName, data.GoodPicture, data.GoodDescribe, data.GoodPrice, data.CurrencyType, data.Id)
	return err
}

func (m *defaultManageGoodsModel) tableName() string {
	return m.table
}
