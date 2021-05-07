package bitpos

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type bitpos struct {
	rdb *redis.Client
}

func (f *bitpos) Args() int           { return -1 }
func (f *bitpos) Deterministic() bool { return false }
func (f *bitpos) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		key string
		bit int64
		r   []int64
	)

	switch len(values) {
	case 0:
		ctx.ResultError(fmt.Errorf("must supply a key"))
		return
	case 1:
		ctx.ResultError(fmt.Errorf("must supply a bit"))
		return
	default:
		key = values[0].Text()
		bit = values[1].Int64()
		r = make([]int64, len(values[2:]))
		for i, v := range values[2:] {
			r[i] = v.Int64()
		}
	}

	res, err := f.rdb.BitPos(context.TODO(), key, bit, r...).Result()
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
	return &bitpos{rdb: rdb}
}
