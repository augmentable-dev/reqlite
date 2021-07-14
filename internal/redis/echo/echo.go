package echo

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type echo struct {
	rdb *redis.Client
}

func (f *echo) Args() int           { return -1 }
func (f *echo) Deterministic() bool { return false }
func (f *echo) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var input string

	if len(values) >= 1 {
		input = values[0].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis echo command"))
		return
	}

	result, err := f.rdb.Echo(context.TODO(), input).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			ctx.ResultNull()
			return
		}
		ctx.ResultError(err)
		return
	}

	ctx.ResultText(fmt.Sprintf("%+q", result))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &echo{rdb}
}
