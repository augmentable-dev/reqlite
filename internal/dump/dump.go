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
	var (
		key    string
		result *redis.StringCmd
		err    error
	)
	if len(values) == 1 {
		key = values[0].Text()

	} else {
		err = fmt.Errorf("must input single arg to dump")
		ctx.ResultError(err)
		return
	}
	result = f.rdb.Dump(context.TODO(), key)

	ctx.ResultText(fmt.Sprint(result.String(), result.FullName(), result.Name()))

}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &dump{rdb}
}
