package get_cmd_test

import (
	"strings"
	"testing"

	get "github.com/augmentable-dev/internal/redis/get_cmd"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestGet(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("get", get.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	mock.ExpectGet("mykey").SetVal("value")
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT get('mykey')")
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

	row = db.QueryRow("SELECT get()")
	err = row.Err()
	if err != nil {
		t.Fatal(err)
	}

	err = row.Scan(&s)
	if strings.Compare(s, "must supply argument to redis get command") != 0 {
		t.Fatal("no err returned with no args supplied")
	}
}
