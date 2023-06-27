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
	yesterdayRewardDetailsFieldNames          = builder.RawFieldNames(&YesterdayRewardDetails{})
	yesterdayRewardDetailsRows                = strings.Join(yesterdayRewardDetailsFieldNames, ",")
	yesterdayRewardDetailsRowsExpectAutoSet   = strings.Join(stringx.Remove(yesterdayRewardDetailsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	yesterdayRewardDetailsRowsWithPlaceHolder = strings.Join(stringx.Remove(yesterdayRewardDetailsFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	yesterdayRewardDetailsModel interface {
		Insert(ctx context.Context, data *YesterdayRewardDetails) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*YesterdayRewardDetails, error)
		Update(ctx context.Context, data *YesterdayRewardDetails) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultYesterdayRewardDetailsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	YesterdayRewardDetails struct {
		UserId       int64     `db:"user_id"`
		RewardType   string    `db:"rewardType"`
		RewardAmount float64   `db:"rewardAmount"`
		RewardTime   time.Time `db:"rewardTime"`
	}
)

func newYesterdayRewardDetailsModel(conn sqlx.SqlConn) *defaultYesterdayRewardDetailsModel {
	return &defaultYesterdayRewardDetailsModel{
		conn:  conn,
		table: "`yesterday_reward_details`",
	}
}

func (m *defaultYesterdayRewardDetailsModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultYesterdayRewardDetailsModel) FindOne(ctx context.Context, userId int64) (*YesterdayRewardDetails, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", yesterdayRewardDetailsRows, m.table)
	var resp YesterdayRewardDetails
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

func (m *defaultYesterdayRewardDetailsModel) Insert(ctx context.Context, data *YesterdayRewardDetails) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, yesterdayRewardDetailsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.RewardType, data.RewardAmount, data.RewardTime)
	return ret, err
}

func (m *defaultYesterdayRewardDetailsModel) Update(ctx context.Context, data *YesterdayRewardDetails) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, yesterdayRewardDetailsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.RewardType, data.RewardAmount, data.RewardTime, data.UserId)
	return err
}

func (m *defaultYesterdayRewardDetailsModel) tableName() string {
	return m.table
}
