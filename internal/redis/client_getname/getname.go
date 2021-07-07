package getname

import (
	"context"
	"errors"
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

	result, err := f.rdb.ClientGetName(context.TODO()).Result()
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
	return &dump{rdb}
}
