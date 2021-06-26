package llen

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type llen struct {
	rdb *redis.Client
}

func (f *llen) Args() int           { return 1 }
func (f *llen) Deterministic() bool { return false }
func (f *llen) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	key := values[0].Text()
	res, err := f.rdb.LLen(context.TODO(), key).Result()
	if err != nil {
		ctx.ResultError(err)
		return
	}

	ctx.ResultInt64(res)
}

func New(rdb *redis.Client) sqlite.Function {
	return &llen{rdb: rdb}
}
