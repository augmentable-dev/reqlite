package hkeys

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type hkeys struct {
	rdb *redis.Client
}

func (f *hkeys) Args() int           { return -1 }
func (f *hkeys) Deterministic() bool { return false }
func (f *hkeys) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var key string

	if len(values) >= 1 {
		key = values[0].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis hkeys command"))
		return
	}

	result := f.rdb.HKeys(context.TODO(), key)

	ctx.ResultText(fmt.Sprintf("%+q", result))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &hkeys{rdb}
}
