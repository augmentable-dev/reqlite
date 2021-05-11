package config_get

import (
	"context"
	"fmt"
	"io"

	"github.com/augmentable-dev/vtab"
	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

const (
	columnKey = iota
	columnVal
	columnPattern
)

var cols = []vtab.Column{
	{Name: "key", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: false, Filters: nil},
	{Name: "val", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: false, Filters: nil},
	{Name: "pattern", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: true, Filters: []sqlite.ConstraintOp{sqlite.INDEX_CONSTRAINT_EQ}},
}

type kv struct {
	key string
	val string
}

type iter struct {
	pattern string
	index   int
	results []kv
}

func newIter(rdb *redis.Client, pattern string) (*iter, error) {
	res, err := rdb.ConfigGet(context.TODO(), pattern).Result()
	if err != nil {
		return nil, err
	}
	results := make([]kv, 0, int(len(res))/2)
	for i := 0; i < len(res); i += 2 {
		results = append(results, kv{key: res[i].(string), val: res[i+1].(string)})
	}

	return &iter{pattern, -1, results}, nil
}

func (i *iter) Column(c int) (interface{}, error) {
	switch c {
	case columnKey:
		return i.results[i.index].key, nil
	case columnVal:
		return i.results[i.index].val, nil
	case columnPattern:
		return i.pattern, nil
	}

	return nil, fmt.Errorf("unknown column")
}

func (i *iter) Next() (vtab.Row, error) {
	i.index++
	if i.index >= len(i.results) {
		return nil, io.EOF
	}
	return i, nil
}

// New returns a config_get virtual table
func New(rdb *redis.Client) sqlite.Module {
	return vtab.NewTableFunc("config_get", cols, func(constraints []vtab.Constraint, order []*sqlite.OrderBy) (vtab.Iterator, error) {
		pattern := "*"
		for _, constraint := range constraints {
			if constraint.Op == sqlite.INDEX_CONSTRAINT_EQ {
				switch constraint.ColIndex {
				case columnPattern:
					pattern = constraint.Value.Text()
				}
			}
		}

		iter, err := newIter(rdb, pattern)
		if err != nil {
			return nil, err
		}

		return iter, nil
	})
}
