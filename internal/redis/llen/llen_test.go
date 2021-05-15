package llen_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/llen"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestLLenOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := int64(100)
	mock.ExpectLLen("test-list").SetVal(want)
	function := llen.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("llen", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT LLEN('test-list')")
	err = row.Err()
	if err != nil {
		t.Fatal(err)
	}

	var res int64
	err = row.Scan(&res)
	if err != nil {
		t.Fatal(err)
	}

	if diff := deep.Equal(res, want); diff != nil {
		t.Error(diff)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestLLenNoArgFail(t *testing.T) {
	function := llen.New(nil)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("llen", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT LLEN()")
	err = row.Err()
	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "wrong number of arguments to function LLEN()" {
		t.Fatal("unexpected error string")
	}

}
