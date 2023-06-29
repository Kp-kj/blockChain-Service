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
	userCryptominerFieldNames          = builder.RawFieldNames(&UserCryptominer{})
	userCryptominerRows                = strings.Join(userCryptominerFieldNames, ",")
	userCryptominerRowsExpectAutoSet   = strings.Join(stringx.Remove(userCryptominerFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userCryptominerRowsWithPlaceHolder = strings.Join(stringx.Remove(userCryptominerFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userCryptominerModel interface {
		Insert(ctx context.Context, data *UserCryptominer) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserCryptominer, error)
		FindOneByUserId(ctx context.Context, userId int64) (*UserCryptominer, error)
		Update(ctx context.Context, data *UserCryptominer) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserCryptominerModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserCryptominer struct {
		UserId            int64        `db:"user_id"`
		CreatedAt         time.Time    `db:"created_at"`
		UpdatedAt         sql.NullTime `db:"updated_at"`
		DeletedAt         sql.NullTime `db:"deleted_at"`
		CryptominerAmount int64        `db:"cryptominer_amount"`
		CryptominerTypeid int64        `db:"cryptominer_typeid"`
	}
)

func newUserCryptominerModel(conn sqlx.SqlConn) *defaultUserCryptominerModel {
	return &defaultUserCryptominerModel{
		conn:  conn,
		table: "`user_cryptominer`",
	}
}

func (m *defaultUserCryptominerModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserCryptominerModel) FindOne(ctx context.Context, userId int64) (*UserCryptominer, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userCryptominerRows, m.table)
	var resp UserCryptominer
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

func (m *defaultUserCryptominerModel) FindOneByUserId(ctx context.Context, userId int64) (*UserCryptominer, error) {
	var resp UserCryptominer
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userCryptominerRows, m.table)
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

func (m *defaultUserCryptominerModel) Insert(ctx context.Context, data *UserCryptominer) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userCryptominerRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.DeletedAt, data.CryptominerAmount, data.CryptominerTypeid)
	return ret, err
}

func (m *defaultUserCryptominerModel) Update(ctx context.Context, newData *UserCryptominer) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userCryptominerRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.CryptominerAmount, newData.CryptominerTypeid, newData.UserId)
	return err
}

func (m *defaultUserCryptominerModel) tableName() string {
	return m.table
}