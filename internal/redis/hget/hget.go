package hget

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type hget struct {
	rdb *redis.Client
}

func (f *hget) Args() int           { return -1 }
func (f *hget) Deterministic() bool { return false }
func (f *hget) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		key   string
		field string
	)

	if len(values) >= 2 {
		key = values[0].Text()
		field = values[1].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis hget command (key,field) "))
		return
	}

	result := f.rdb.HGet(context.TODO(), key, field)

	ctx.ResultText(fmt.Sprintf("%+q", result))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &hget{rdb}
}
