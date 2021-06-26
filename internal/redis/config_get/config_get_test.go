package config_get_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/config_get"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

type kv struct {
	Key string
	Val string
}

func TestConfigGetOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := []interface{}{"key1", "val1", "key2", "val2"}
	wantRows := []kv{{"key1", "val1"}, {"key2", "val2"}}
	mock.ExpectConfigGet("*max-*-entries*").SetVal(want)
	mod := config_get.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateModule("config_get", mod,
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

	var rows []kv
	err = db.Select(&rows, "SELECT * FROM CONFIG_GET('*max-*-entries*')")
	if err != nil {
		t.Fatal(err)
	}

	if diff := deep.Equal(rows, wantRows); diff != nil {
		t.Error(diff)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
