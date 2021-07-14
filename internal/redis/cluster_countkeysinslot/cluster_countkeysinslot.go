package cluster_countkeysinslot

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type clustercountkeysinslot struct {
	rdb *redis.Client
}

func (f *clustercountkeysinslot) Args() int           { return -1 }
func (f *clustercountkeysinslot) Deterministic() bool { return false }
func (f *clustercountkeysinslot) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	if len(values) < 1 {
		ctx.ResultError(fmt.Errorf("must supply a slot (type int)"))
		return
	}
	slot := values[0].Int()
	id, err := f.rdb.ClusterCountKeysInSlot(context.TODO(), slot).Result()
	if err != nil {
		ctx.ResultError(err)
		return
	}
	ctx.ResultInt64(id)
}

func New(rdb *redis.Client) sqlite.Function {
	return &clustercountkeysinslot{rdb: rdb}
}
