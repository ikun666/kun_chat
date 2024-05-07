package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		withSession(session sqlx.Session) GroupModel
		GetGroups(ctx context.Context, uid int64) ([]group, error)
		DeleteByUidName(ctx context.Context, uid int64, name string) error
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn),
	}
}

func (m *customGroupModel) withSession(session sqlx.Session) GroupModel {
	return NewGroupModel(sqlx.NewSqlConnFromSession(session))
}
func (m *customGroupModel) DeleteByUidName(ctx context.Context, uid int64, name string) error {
	query := fmt.Sprintf("delete from %s where `owner_id` = ? and `name` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, uid, name)
	return err
}

func (m *customGroupModel) GetGroups(ctx context.Context, uid int64) ([]group, error) {
	query := fmt.Sprintf("select `id`,`name` from %s where `owner_id` = ? ", m.table)
	var resp []group
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
