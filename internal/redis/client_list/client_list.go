package client_list

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type clientlist struct {
	rdb *redis.Client
}

func (f *clientlist) Args() int           { return 0 }
func (f *clientlist) Deterministic() bool { return false }
func (f *clientlist) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	id, err := f.rdb.ClientList(context.TODO()).Result()
	if err != nil {
		ctx.ResultError(err)
		return
	}

	ctx.ResultText(id)
}

func New(rdb *redis.Client) sqlite.Function {
	return &clientlist{rdb: rdb}
}
