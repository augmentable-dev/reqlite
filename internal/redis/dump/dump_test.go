package dump

import (
	"fmt"
	"testing"

	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestDump(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	expected := "\x00\xc0\n\t\x00\xbem\x06\x89Z(\x00\n"

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("dump", New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	mock.ExpectDump("mykey").SetVal(expected)
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT dump('mykey')")
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
func TestDumpErr(t *testing.T) {
	rdb, _ := redismock.NewClientMock()

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("dump", New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT dump()")
	var s string
	err = row.Scan(&s)
	if err == nil {
		t.Fatal(fmt.Errorf("dump returning no error without args"))
	}

}
func TestDumpNil(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("dump", New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})
	mock.ExpectDump("nokey")
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT dump('nokey')")
	var s string
	err = row.Scan(&s)
	if s != "" {
		t.Fatalf("string value returned from uninitialized key %s", s)
	}
	if err == nil {
		t.Fatal("no error from uninitialized key")
	}

}
