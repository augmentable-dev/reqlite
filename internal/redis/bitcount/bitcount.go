package bitcount

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type bitcount struct {
	rdb *redis.Client
}

func (f *bitcount) Args() int           { return -1 }
func (f *bitcount) Deterministic() bool { return false }
func (f *bitcount) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		key   string
		start int64
		end   int64
	)

	switch len(values) {
	case 0:
		ctx.ResultError(fmt.Errorf("must supply a key"))
		return
	case 1:
		key = values[0].Text()
	case 2:
		ctx.ResultError(fmt.Errorf("must supply a start and end"))
		return
	case 3:
		key = values[0].Text()
		start = values[1].Int64()
		end = values[2].Int64()
	}

	res, err := f.rdb.BitCount(context.TODO(), key, &redis.BitCount{Start: start, End: end}).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			ctx.ResultNull()
			return
		}

		ctx.ResultError(err)
		return
	}

	ctx.ResultInt64(res)
}

func New(rdb *redis.Client) sqlite.Function {
	return &bitcount{rdb: rdb}
}
