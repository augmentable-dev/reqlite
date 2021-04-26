package hgetall

import (
	"context"
	"fmt"
	"io"

	"github.com/augmentable-dev/vtab"
	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

const (
	columnHKey = iota
	columnHVal
	columnKey
)

var cols = []vtab.Column{
	{Name: "hkey", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: false, Filters: nil},
	{Name: "hval", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: false, Filters: nil},
	{Name: "key", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: true, Filters: []sqlite.ConstraintOp{sqlite.INDEX_CONSTRAINT_EQ}},
}

type kv struct {
	key string
	val string
}

type iter struct {
	key     string
	index   int
	results []kv
}

func newIter(key string) (*iter, error) {
	rdb := redis.NewClient(&redis.Options{})
	h, err := rdb.HGetAll(context.TODO(), key).Result()
	if err != nil {
		return nil, err
	}
	res := make([]kv, 0, len(h))
	for k, v := range h {
		res = append(res, kv{k, v})
	}
	return &iter{key, -1, res}, nil
}

func (i *iter) Column(c int) (interface{}, error) {
	switch c {
	case columnHKey:
		return i.results[i.index].key, nil
	case columnHVal:
		return i.results[i.index].val, nil
	case columnKey:
		return i.key, nil
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

// New returns an lrange virtual table
func New() sqlite.Module {
	return vtab.NewTableFunc("hgetall", cols, func(constraints []vtab.Constraint, order []*sqlite.OrderBy) (vtab.Iterator, error) {
		var key string
		for _, constraint := range constraints {
			if constraint.Op == sqlite.INDEX_CONSTRAINT_EQ {
				switch constraint.ColIndex {
				case columnKey:
					key = constraint.Value.Text()
				}
			}
		}

		iter, err := newIter(key)
		if err != nil {
			return nil, err
		}

		return iter, nil
	})
}
