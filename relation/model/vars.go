package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type group struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
