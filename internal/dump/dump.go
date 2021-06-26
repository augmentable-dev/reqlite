package dump

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type dump struct {
	rdb *redis.Client
}

func (f *dump) Args() int           { return -1 }
func (f *dump) Deterministic() bool { return false }
func (f *dump) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var key string

	if len(values) >= 1 {
		key = values[0].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis dump command"))
		return
	}

	result, err := f.rdb.Dump(context.TODO(), key).Result()
	if err != nil {
		ctx.ResultError(err)
		return
	}

	ctx.ResultText(result)

}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &dump{rdb}
}
