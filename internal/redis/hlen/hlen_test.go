package hlen_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/hlen"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestHLen(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("hlen", hlen.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	mock.ExpectHLen("mykey").SetVal(1024)
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT hlen('mykey')")
	err = row.Err()
	if err != nil {
		t.Fatal(err)
	}

	var s string
	err = row.Scan(&s)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
