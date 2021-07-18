package hlen

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type hlen struct {
	rdb *redis.Client
}

func (f *hlen) Args() int           { return -1 }
func (f *hlen) Deterministic() bool { return false }
func (f *hlen) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var key string

	if len(values) >= 1 {
		key = values[0].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis hlen command"))
		return
	}

	result := f.rdb.HLen(context.TODO(), key)

	ctx.ResultInt64(result.Val())
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &hlen{rdb}
}
