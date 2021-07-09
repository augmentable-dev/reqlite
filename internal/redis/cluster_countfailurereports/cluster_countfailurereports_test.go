package cluster_countfailurereports_test

import (
	"fmt"
	"testing"

	"github.com/augmentable-dev/reqlite/internal/redis/cluster_countfailurereports"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestCluster_countfailurereports_OK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := int64(1024)
	mock.ExpectClusterCountFailureReports("abcdefghijklmnopqrstuvwxyz0123456789aaab").SetVal(want)
	function := cluster_countfailurereports.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("clustercountfailurereports", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT clustercountfailurereports('abcdefghijklmnopqrstuvwxyz0123456789aaab')")
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
func TestCluster_countfailurereports_extra_argsOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := int64(1024)
	mock.ExpectClusterCountFailureReports("abcdefghijklmnopqrstuvwxyz0123456789aaab").SetVal(want)
	function := cluster_countfailurereports.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("clustercountfailurereports", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT clustercountfailurereports('abcdefghijklmnopqrstuvwxyz0123456789aaab','this','should','be','fine')")
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
func TestCluster_countfailurereports_error(t *testing.T) {
	rdb, _ := redismock.NewClientMock()

	function := cluster_countfailurereports.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("clustercountfailurereports", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT clustercountfailurereports()")
	var s string
	err = row.Scan(&s)
	if err == nil {
		t.Fatal(fmt.Errorf("clustercountfailurereports given no input and did not return err, returned : %s", s))
	}
}
