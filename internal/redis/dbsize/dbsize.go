package dbsize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type dbsize struct {
	rdb *redis.Client
}

func (f *dbsize) Args() int           { return 0 }
func (f *dbsize) Deterministic() bool { return false }
func (f *dbsize) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	res, err := f.rdb.DBSize(context.TODO()).Result()
	if err != nil {
		ctx.ResultError(err)
		return
	}

	ctx.ResultInt64(res)
}

func New(rdb *redis.Client) sqlite.Function {
	return &dbsize{rdb: rdb}
}
