package getrange

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type getrange struct {
	rdb *redis.Client
}

func (f *getrange) Args() int           { return -1 }
func (f *getrange) Deterministic() bool { return false }
func (f *getrange) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var (
		key   string
		start int64
		end   int64
	)

	if len(values) >= 3 {
		key = values[0].Text()
		start = values[1].Int64()
		end = values[2].Int64()
	} else {
		ctx.ResultError(fmt.Errorf("must supply (key,start,end) to getrange"))
		return
	}

	result := f.rdb.GetRange(context.TODO(), key, start, end)

	ctx.ResultText(fmt.Sprintf("%+q", result))
}

// New returns a sqlite function for reading the contents of a file
func New(rdb *redis.Client) sqlite.Function {
	return &getrange{rdb}
}
