package clientid

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type clientid struct {
	rdb *redis.Client
}

func (f *clientid) Args() int           { return 0 }
func (f *clientid) Deterministic() bool { return false }
func (f *clientid) Apply(ctx *sqlite.Context, values ...sqlite.Value) {

	id, err := f.rdb.ClientID(context.TODO()).Result()
	if err != nil {
		ctx.ResultError(err)
		return
	}

	ctx.ResultInt64(id)
}

func New(rdb *redis.Client) sqlite.Function {
	return &clientid{rdb: rdb}
}
