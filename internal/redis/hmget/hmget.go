package hmget

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type hmget struct {
	rdb *redis.Client
}

func (f *hmget) Args() int           { return -1 }
func (f *hmget) Deterministic() bool { return false }
func (f *hmget) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		key    string
		fields []string
	)

	if len(values) >= 2 {
		key = values[0].Text()
		for i := 1; i < len(values); i++ {
			fields = append(fields, values[i].Text())
		}
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis hmget command (key,field1,...,fieldn"))
		return
	}

	result := f.rdb.HMGet(context.TODO(), key, fields...)

	ctx.ResultText(fmt.Sprintf("%+q", result.String()))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &hmget{rdb}
}
