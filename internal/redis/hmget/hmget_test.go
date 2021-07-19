package hmget_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/hmget"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestHmget(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	fields := []string{"field1", "field2", "field3"}
	values := []interface{}{"val1", "val2", "val3"}
	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("hmget", hmget.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})
	mock.ExpectHMGet("mykey", fields...).SetVal(values)
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT hmget('mykey','field1','field2','field3')")
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
