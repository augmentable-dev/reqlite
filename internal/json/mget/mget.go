package mget

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/augmentable-dev/vtab"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"go.riyazali.net/sqlite"
)

const (
	columnJSON = iota
	columnPath
	columnKeys
)

var cols = []vtab.Column{
	{Name: "json", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: false, Filters: nil},
	{Name: "path", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: true, Filters: []sqlite.ConstraintOp{sqlite.INDEX_CONSTRAINT_EQ}},
	{Name: "keys", Type: sqlite.SQLITE_TEXT, NotNull: false, Hidden: true, Filters: []sqlite.ConstraintOp{sqlite.INDEX_CONSTRAINT_EQ}},
}

type iter struct {
	keys    []string
	path    string
	index   int
	results []string
}

func newIter(rdb *redis.Client, keys, path string) (*iter, error) {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(rdb)

	r := csv.NewReader(strings.NewReader(keys))
	k, err := r.Read()
	if err != nil {
		return nil, err
	}

	resRaw, err := rh.JSONMGet(path, k...)
	if err != nil {
		return nil, err
	}

	res := resRaw.([]interface{})
	out := make([]string, len(res))
	for i, r := range res {
		out[i] = string(r.([]byte))
	}

	return &iter{k, path, -1, out}, nil
}

func (i *iter) Column(c int) (interface{}, error) {
	switch c {
	case columnJSON:
		return i.results[i.index], nil
	case columnPath:
		return i.path, nil
	case columnKeys:
		return strings.Join(i.keys, ","), nil
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

// New returns a json mget virtual table
func New(rdb *redis.Client) sqlite.Module {
	return vtab.NewTableFunc("json_mget", cols, func(constraints []vtab.Constraint, order []*sqlite.OrderBy) (vtab.Iterator, error) {
		var keys, path string
		path = "."
		for _, constraint := range constraints {
			if constraint.Op == sqlite.INDEX_CONSTRAINT_EQ {
				switch constraint.ColIndex {
				case columnKeys:
					keys = constraint.Value.Text()
				case columnPath:
					path = constraint.Value.Text()
				}
			}
		}

		iter, err := newIter(rdb, keys, path)
		if err != nil {
			return nil, err
		}

		return iter, nil
	})
}
