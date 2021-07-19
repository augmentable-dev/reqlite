package hvals

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type hvals struct {
	rdb *redis.Client
}

func (f *hvals) Args() int           { return -1 }
func (f *hvals) Deterministic() bool { return false }
func (f *hvals) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var key string

	if len(values) >= 1 {
		key = values[0].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis hvals command (key)"))
		return
	}

	result := f.rdb.HVals(context.TODO(), key)

	ctx.ResultText(fmt.Sprintf("%+q", result.String()))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &hvals{rdb}
}
