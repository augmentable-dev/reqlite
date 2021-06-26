package get

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"go.riyazali.net/sqlite"
)

type get struct {
	rh *rejson.Handler
}

func (f *get) Args() int           { return -1 }
func (f *get) Deterministic() bool { return false }
func (f *get) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	var key, path string

	switch len(values) {
	case 0:
		ctx.ResultError(fmt.Errorf("must supply a key"))
		return
	case 1:
		key = values[0].Text()
		path = "."
	case 2:
		fallthrough
	default:
		key = values[0].Text()
		path = values[1].Text()
	}

	res, err := f.rh.JSONGet(key, path)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			ctx.ResultNull()
			return
		}

		ctx.ResultError(err)
		return
	}

	ctx.ResultText(string(res.([]byte)))
}

func New(rdb *redis.Client) sqlite.Function {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(rdb)
	return &get{rh: rh}
}
