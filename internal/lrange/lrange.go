package lrange

import (
	"context"
	"fmt"
	"io"

	"github.com/augmentable-dev/vtab"
	"github.com/go-redis/redis/v8"
	"go.riyazali.net/sqlite"
)

const (
	columnValue = iota
	columnKey
	columnStart
	columnStop
)

var cols = []vtab.Column{
	{Name: "value", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: false, Filters: nil},
	{Name: "key", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: true, Filters: []sqlite.ConstraintOp{sqlite.INDEX_CONSTRAINT_EQ}},
	{Name: "start", Type: sqlite.SQLITE_INTEGER, NotNull: false, Hidden: true, Filters: []sqlite.ConstraintOp{sqlite.INDEX_CONSTRAINT_EQ}},
	{Name: "stop", Type: sqlite.SQLITE_INTEGER, NotNull: false, Hidden: true, Filters: []sqlite.ConstraintOp{sqlite.INDEX_CONSTRAINT_EQ}},
}

type iter struct {
	key     string
	start   int
	stop    int
	index   int
	results []string
}

func newIter(key string, start, stop int) (*iter, error) {
	rdb := redis.NewClient(&redis.Options{})
	res, err := rdb.LRange(context.TODO(), key, int64(start), int64(stop)).Result()
	if err != nil {
		return nil, err
	}
	return &iter{key, start, stop, -1, res}, nil
}

func (i *iter) Column(c int) (interface{}, error) {
	switch c {
	case columnValue:
		return i.results[i.index], nil
	case columnKey:
		return i.key, nil
	case columnStart:
		return i.start, nil
	case columnStop:
		return i.stop, nil
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
	return vtab.NewTableFunc("lrange", cols, func(constraints []vtab.Constraint, order []*sqlite.OrderBy) (vtab.Iterator, error) {
		var (
			key     string
			start   int
			stop    int
			stopSet bool
		)
		for _, constraint := range constraints {
			if constraint.Op == sqlite.INDEX_CONSTRAINT_EQ {
				switch constraint.ColIndex {
				case columnKey:
					key = constraint.Value.Text()
				case columnStart:
					start = constraint.Value.Int()
				case columnStop:
					stop = constraint.Value.Int()
					stopSet = true
				}
			}
		}

		// if stop is not set, make it the same as start
		// so that SELECT * FROM lrange('key', 10) => LRANGE key 10 10
		if !stopSet {
			stop = start
		}

		iter, err := newIter(key, start, stop)
		if err != nil {
			return nil, err
		}

		return iter, nil
	})
}
