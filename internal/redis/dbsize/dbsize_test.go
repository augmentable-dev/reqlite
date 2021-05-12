package dbsize_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/dbsize"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestDBSizeOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := int64(100)
	mock.ExpectDBSize().SetVal(want)
	function := dbsize.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("dbsize", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT DBSIZE()")
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

// TODO add test cases for non-happy paths
