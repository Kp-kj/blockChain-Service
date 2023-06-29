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
	propFieldNames          = builder.RawFieldNames(&Prop{})
	propRows                = strings.Join(propFieldNames, ",")
	propRowsExpectAutoSet   = strings.Join(stringx.Remove(propFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	propRowsWithPlaceHolder = strings.Join(stringx.Remove(propFieldNames, "`prop_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	propModel interface {
		Insert(ctx context.Context, data *Prop) (sql.Result, error)
		FindOne(ctx context.Context, propId int64) (*Prop, error)
		FindOneByPropId(ctx context.Context, propId int64) (*Prop, error)
		Update(ctx context.Context, data *Prop) error
		Delete(ctx context.Context, propId int64) error
	}

	defaultPropModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Prop struct {
		PropId       int64          `db:"prop_id"`
		CreatedAt    time.Time      `db:"created_at"`
		UpdatedAt    sql.NullTime   `db:"updated_at"`
		DeletedAt    sql.NullTime   `db:"deleted_at"`
		UserId       sql.NullInt64  `db:"user_id"`
		PropTypeid   int64          `db:"prop_typeid"`
		PropName     string         `db:"prop_name"`
		PropPicture  sql.NullString `db:"prop_picture"`
		PropPrice    int64          `db:"prop_price"`
		PropDescribe sql.NullString `db:"prop_describe"`
		PurchaseTime sql.NullTime   `db:"purchase_time"`
	}
)

func newPropModel(conn sqlx.SqlConn) *defaultPropModel {
	return &defaultPropModel{
		conn:  conn,
		table: "`prop`",
	}
}

func (m *defaultPropModel) Delete(ctx context.Context, propId int64) error {
	query := fmt.Sprintf("delete from %s where `prop_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, propId)
	return err
}

func (m *defaultPropModel) FindOne(ctx context.Context, propId int64) (*Prop, error) {
	query := fmt.Sprintf("select %s from %s where `prop_id` = ? limit 1", propRows, m.table)
	var resp Prop
	err := m.conn.QueryRowCtx(ctx, &resp, query, propId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPropModel) FindOneByPropId(ctx context.Context, propId int64) (*Prop, error) {
	var resp Prop
	query := fmt.Sprintf("select %s from %s where `prop_id` = ? limit 1", propRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, propId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPropModel) Insert(ctx context.Context, data *Prop) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, propRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PropId, data.DeletedAt, data.UserId, data.PropTypeid, data.PropName, data.PropPicture, data.PropPrice, data.PropDescribe, data.PurchaseTime)
	return ret, err
}

func (m *defaultPropModel) Update(ctx context.Context, newData *Prop) error {
	query := fmt.Sprintf("update %s set %s where `prop_id` = ?", m.table, propRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.UserId, newData.PropTypeid, newData.PropName, newData.PropPicture, newData.PropPrice, newData.PropDescribe, newData.PurchaseTime, newData.PropId)
	return err
}

func (m *defaultPropModel) tableName() string {
	return m.table
}
