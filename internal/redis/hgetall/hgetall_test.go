package hgetall_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/hgetall"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

type kv struct {
	Key string `db:"hkey"`
	Val string `db:"hval"`
}

func TestHGetAllOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := []kv{
		{Key: "one", Val: "1"},
		{Key: "two", Val: "2"},
		{Key: "three", Val: "3"},
	}
	wantMap := make(map[string]string, len(want))
	for _, pair := range want {
		wantMap[pair.Key] = pair.Val
	}
	mock.ExpectHGetAll("test-hash").SetVal(wantMap)
	mod := hgetall.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateModule("hgetall", mod,
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

	rows := []kv{}
	// the ORDER BY in the query is necessary since there's no guarantee what order the key/value pairs
	// will be returned in, leading to a flaky test case in the deep equal below (where order matters)
	err = db.Select(&rows, "SELECT * FROM HGETALL('test-hash') ORDER BY hval")
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
