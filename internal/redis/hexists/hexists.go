package hexists

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type hexists struct {
	rdb *redis.Client
}

func (f *hexists) Args() int           { return -1 }
func (f *hexists) Deterministic() bool { return false }
func (f *hexists) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		key   string
		field string
	)

	if len(values) >= 2 {
		key = values[0].Text()
		field = values[1].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis hexists command (key,field) "))
		return
	}

	result := f.rdb.HExists(context.TODO(), key, field)

	ctx.ResultText(fmt.Sprintf("%+q", result))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &hexists{rdb}
}
