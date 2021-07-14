package cluster_count_failure_reports

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

type clustercountfailurereports struct {
	rdb *redis.Client
}

func (f *clustercountfailurereports) Args() int           { return -1 }
func (f *clustercountfailurereports) Deterministic() bool { return false }
func (f *clustercountfailurereports) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	if len(values) < 1 {
		ctx.ResultError(fmt.Errorf("must supply a node id"))
		return
	}
	node_id := values[0].Text()
	id, err := f.rdb.ClusterCountFailureReports(context.TODO(), node_id).Result()
	if err != nil {
		ctx.ResultError(err)
		return
	}

	ctx.ResultInt64(id)
}

func New(rdb *redis.Client) sqlite.Function {
	return &clustercountfailurereports{rdb: rdb}
}
