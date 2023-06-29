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
	userPropFieldNames          = builder.RawFieldNames(&UserProp{})
	userPropRows                = strings.Join(userPropFieldNames, ",")
	userPropRowsExpectAutoSet   = strings.Join(stringx.Remove(userPropFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPropRowsWithPlaceHolder = strings.Join(stringx.Remove(userPropFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPropModel interface {
		Insert(ctx context.Context, data *UserProp) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserProp, error)
		FindOneByUserId(ctx context.Context, userId int64) (*UserProp, error)
		Update(ctx context.Context, data *UserProp) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserPropModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserProp struct {
		UserId     int64        `db:"user_id"`
		CreatedAt  time.Time    `db:"created_at"`
		UpdatedAt  sql.NullTime `db:"updated_at"`
		DeletedAt  sql.NullTime `db:"deleted_at"`
		PropAmount int64        `db:"prop_amount"`
		PropTypeid int64        `db:"prop_typeid"`
	}
)

func newUserPropModel(conn sqlx.SqlConn) *defaultUserPropModel {
	return &defaultUserPropModel{
		conn:  conn,
		table: "`user_prop`",
	}
}

func (m *defaultUserPropModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserPropModel) FindOne(ctx context.Context, userId int64) (*UserProp, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userPropRows, m.table)
	var resp UserProp
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserPropModel) FindOneByUserId(ctx context.Context, userId int64) (*UserProp, error) {
	var resp UserProp
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userPropRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserPropModel) Insert(ctx context.Context, data *UserProp) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userPropRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.DeletedAt, data.PropAmount, data.PropTypeid)
	return ret, err
}

func (m *defaultUserPropModel) Update(ctx context.Context, newData *UserProp) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userPropRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.PropAmount, newData.PropTypeid, newData.UserId)
	return err
}

func (m *defaultUserPropModel) tableName() string {
	return m.table
}
