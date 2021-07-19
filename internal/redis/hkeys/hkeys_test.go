package hkeys_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/hkeys"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestDump(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	myFields := []string{"field1", "field2", "field3"}
	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("hkeys", hkeys.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	mock.ExpectHKeys("mykey").SetVal(myFields)
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT hkeys('mykey')")
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
