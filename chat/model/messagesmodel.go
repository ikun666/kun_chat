package model

import (
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessagesModel = (*customMessagesModel)(nil)
var BulkInserter *sqlx.BulkInserter

// var conn sqlx.SqlConn
// blk, err := sqlx.NewBulkInserter(conn, "insert into user (id, name) values (?, ?)")
//
//	if err != nil {
//		panic(err)
//	}
//
// defer blk.Flush()
// blk.Insert(1, "test1")
// blk.Insert(2, "test2")
type (
	// MessagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessagesModel.
	MessagesModel interface {
		messagesModel
		withSession(session sqlx.Session) MessagesModel
		BulkInsert(msgs []Message)
	}

	customMessagesModel struct {
		*defaultMessagesModel
	}
)

// NewMessagesModel returns a model for the database table.
func NewMessagesModel(conn sqlx.SqlConn) MessagesModel {

	m := &customMessagesModel{
		defaultMessagesModel: newMessagesModel(conn),
	}
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, messagesRowsExpectAutoSet)
	// ret, err := m.conn.ExecCtx(ctx, query, data.FormId, data.TargetId, data.Type, data.Media, data.Content, data.Pic, data.Url, data.Desc, data.Amount)
	// return ret, err
	BulkInserter, _ = sqlx.NewBulkInserter(conn, query)
	BulkInserter.SetResultHandler(func(result sql.Result, err error) {
		if err != nil {
			logx.Error(err)
			return
		}
		fmt.Println(result.RowsAffected())
	})
	return m
}

func (m *customMessagesModel) withSession(session sqlx.Session) MessagesModel {
	return NewMessagesModel(sqlx.NewSqlConnFromSession(session))
}
func (m *customMessagesModel) BulkInsert(msgs []Message) {
	for i := range msgs {
		BulkInserter.Insert(msgs[i].FormId, msgs[i].TargetId, msgs[i].Type, msgs[i].Media, msgs[i].Content, msgs[i].Pic, msgs[i].Url, msgs[i].Desc, msgs[i].Amount)
	}
}
