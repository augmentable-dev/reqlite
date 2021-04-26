package hgetall_test

import (
	"database/sql"
	"testing"

	"github.com/augmentable-dev/reqlite/internal/hgetall"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestHGetAll(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	mock.ExpectHGetAll("test-hash").SetVal(map[string]string{
		"one": "1",
		"two": "2",
	})

	mod := hgetall.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateModule("hgetall", mod,
			sqlite.EponymousOnly(true),
			sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM HGETALL('test-hash')")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	// TODO this maybe could be simplified with: https://github.com/jmoiron/sqlx
	_, contents, err := GetContents(rows)
	if err != nil {
		t.Fatal(err)
	}

	if len(contents) != 2 {
		t.Fail()
	}

	if contents[0][0] != "one" || contents[0][1] != "1" {
		t.Fail()
	}

	if contents[1][0] != "two" || contents[1][1] != "2" {
		t.Fail()
	}

	err = rows.Err()
	if err != nil {
		t.Fatal(err)
	}
}

func GetContents(rows *sql.Rows) (int, [][]string, error) {
	count := 0
	columns, err := rows.Columns()
	if err != nil {
		return count, nil, err
	}

	pointers := make([]interface{}, len(columns))
	container := make([]sql.NullString, len(columns))
	var ret [][]string

	for i := range pointers {
		pointers[i] = &container[i]
	}

	for rows.Next() {
		err = rows.Scan(pointers...)
		if err != nil {
			return count, nil, err
		}

		r := make([]string, len(columns))
		for i, c := range container {
			if c.Valid {
				r[i] = c.String
			} else {
				r[i] = "NULL"
			}
		}
		ret = append(ret, r)
	}
	return count, ret, err

}
