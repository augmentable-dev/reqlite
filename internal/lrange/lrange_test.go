package lrange_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/lrange"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestLRangeOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := []string{"one", "two", "three"}
	mock.ExpectLRange("test-list", 0, 5).SetVal(want)
	mod := lrange.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateModule("lrange", mod,
			sqlite.EponymousOnly(true),
			sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := []string{}
	err = db.Select(&rows, "SELECT * FROM LRANGE('test-list', 0, 5)")
	if err != nil {
		t.Fatal(err)
	}

	if diff := deep.Equal(rows, want); diff != nil {
		t.Error(diff)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
