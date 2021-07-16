package getrange_test

import (
	"testing"

	getrange "github.com/augmentable-dev/reqlite/internal/redis/getrange"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestGetrange(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("getrange", getrange.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	mock.ExpectGetRange("my-key", -4, -1).SetVal("key")
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT getrange('my-key',-4,-1)")
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
func TestGetrangeErr(t *testing.T) {
	rdb, _ := redismock.NewClientMock()

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("getrange", getrange.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var s string
	row := db.QueryRow("SELECT getrange()")
	err = row.Scan(&s)
	if err == nil && s != "must supply (key,start,end) to getrange" {
		t.Fatal("no error returned when no input provided")
	}

}
