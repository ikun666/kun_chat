package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelationModel = (*customRelationModel)(nil)

type (
	// RelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationModel.
	RelationModel interface {
		relationModel
		withSession(session sqlx.Session) RelationModel
		AddFriend(ctx context.Context, uid, tid, tp int64) error
		DelFriend(ctx context.Context, uid, tid, tp int64) error
		GetFriends(ctx context.Context, uid, tp int64) ([]int64, error)
		FindFriendByUidTidType(ctx context.Context, uid, tid, tp int64) (bool, error)
	}

	customRelationModel struct {
		*defaultRelationModel
	}
)

// NewRelationModel returns a model for the database table.
func NewRelationModel(conn sqlx.SqlConn) RelationModel {
	return &customRelationModel{
		defaultRelationModel: newRelationModel(conn),
	}
}

func (m *customRelationModel) withSession(session sqlx.Session) RelationModel {
	return NewRelationModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customRelationModel) GetFriends(ctx context.Context, uid, tp int64) ([]int64, error) {

	query := fmt.Sprintf("select `target_id` from %s where `owner_id` = ? and `type` = ? ", m.table)
	var resp []int64
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, tp)

	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customRelationModel) FindFriendByUidTidType(ctx context.Context, uid, tid, tp int64) (bool, error) {

	query := fmt.Sprintf("select `id` from %s where `owner_id` = ? and `target_id` = ? and `type` = ? limit 1", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, uid, tid, tp)

	switch err {
	case nil:
		return true, nil
	case sqlx.ErrNotFound:
		return false, ErrNotFound
	default:
		return false, err
	}
}
func (m *customRelationModel) AddFriend(ctx context.Context, uid, tid, tp int64) error {
	err := m.conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, relationRowsExpectAutoSet)
		_, err := session.ExecCtx(ctx, query, uid, tid, tp, "")
		if err != nil {
			return err
		}
		query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, relationRowsExpectAutoSet)
		_, err = session.ExecCtx(ctx, query, tid, uid, tp, "")
		if err != nil {
			return err
		}
		return err
	})
	return err
}
func (m *customRelationModel) DelFriend(ctx context.Context, uid, tid, tp int64) error {
	err := m.conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		query := fmt.Sprintf("delete from %s where `owner_id` = ? and `target_id` = ? and `type` = ?", m.table)
		_, err := session.ExecCtx(ctx, query, uid, tid, tp)
		if err != nil {
			return err
		}
		query = fmt.Sprintf("delete from %s where `owner_id` = ? and `target_id` = ? and `type` = ?", m.table)
		_, err = session.ExecCtx(ctx, query, tid, uid, tp)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
