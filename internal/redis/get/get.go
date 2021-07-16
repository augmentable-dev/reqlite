package get

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type get struct {
	rdb *redis.Client
}

func (f *get) Args() int           { return -1 }
func (f *get) Deterministic() bool { return false }
func (f *get) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var input string

	if len(values) >= 1 {
		input = values[0].Text()
	} else {
		ctx.ResultError(fmt.Errorf("must supply argument to redis get command"))
		return
	}

	result := f.rdb.Get(context.TODO(), input)

	ctx.ResultText(fmt.Sprintf("%+q", result))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &get{rdb}
}
