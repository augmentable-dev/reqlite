package bitcount_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/bitcount"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestBitcountOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := int64(4)
	mock.ExpectBitCount("test-key", &redis.BitCount{Start: 1, End: 1}).SetVal(want)
	function := bitcount.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("bitcount", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT BITCOUNT('test-key', 1, 1)")
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

// TODO add test cases for no start/end and non-happy path
