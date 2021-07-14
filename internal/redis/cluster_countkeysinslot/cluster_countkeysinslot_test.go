package cluster_countkeysinslot_test

import (
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/cluster_countkeysinslot"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestClusterCountKeysInSlotOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := int64(16)
	slot := 1024
	mock.ExpectClusterCountKeysInSlot(slot).SetVal(want)
	function := cluster_countkeysinslot.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("cluster_countkeysinslot", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT cluster_countkeysinslot(1024)")
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
