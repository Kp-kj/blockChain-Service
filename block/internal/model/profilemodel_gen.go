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
	profileFieldNames          = builder.RawFieldNames(&Profile{})
	profileRows                = strings.Join(profileFieldNames, ",")
	profileRowsExpectAutoSet   = strings.Join(stringx.Remove(profileFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	profileRowsWithPlaceHolder = strings.Join(stringx.Remove(profileFieldNames, "`profile_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	profileModel interface {
		Insert(ctx context.Context, data *Profile) (sql.Result, error)
		FindOne(ctx context.Context, profileId int64) (*Profile, error)
		FindOneByProfileId(ctx context.Context, profileId int64) (*Profile, error)
		Update(ctx context.Context, data *Profile) error
		Delete(ctx context.Context, profileId int64) error
	}

	defaultProfileModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Profile struct {
		ProfileId   int64        `db:"profile_id"`
		CreatedAt   time.Time    `db:"created_at"`
		UpdatedAt   sql.NullTime `db:"updated_at"`
		DeletedAt   sql.NullTime `db:"deleted_at"`
		UserId      int64        `db:"user_id"`
		TwitterName string       `db:"twitter_name"`
		UserName    string       `db:"user_name"`
		Avatar      string       `db:"avatar"`
		IsNew       int64        `db:"is_new"`
	}
)

func newProfileModel(conn sqlx.SqlConn) *defaultProfileModel {
	return &defaultProfileModel{
		conn:  conn,
		table: "`profile`",
	}
}

func (m *defaultProfileModel) withSession(session sqlx.Session) *defaultProfileModel {
	return &defaultProfileModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`profile`",
	}
}

func (m *defaultProfileModel) Delete(ctx context.Context, profileId int64) error {
	query := fmt.Sprintf("delete from %s where `profile_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, profileId)
	return err
}

func (m *defaultProfileModel) FindOne(ctx context.Context, profileId int64) (*Profile, error) {
	query := fmt.Sprintf("select %s from %s where `profile_id` = ? limit 1", profileRows, m.table)
	var resp Profile
	err := m.conn.QueryRowCtx(ctx, &resp, query, profileId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProfileModel) FindOneByProfileId(ctx context.Context, profileId int64) (*Profile, error) {
	var resp Profile
	query := fmt.Sprintf("select %s from %s where `profile_id` = ? limit 1", profileRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, profileId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProfileModel) Insert(ctx context.Context, data *Profile) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, profileRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProfileId, data.DeletedAt, data.UserId, data.TwitterName, data.UserName, data.Avatar, data.IsNew)
	return ret, err
}

func (m *defaultProfileModel) Update(ctx context.Context, newData *Profile) error {
	query := fmt.Sprintf("update %s set %s where `profile_id` = ?", m.table, profileRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.UserId, newData.TwitterName, newData.UserName, newData.Avatar, newData.IsNew, newData.ProfileId)
	return err
}

func (m *defaultProfileModel) tableName() string {
	return m.table
}
