package getname_test

import (
	"fmt"
	"strings"
	"testing"

	getname "github.com/augmentable-dev/reqlite/internal/redis/client_getname"
	_ "github.com/augmentable-dev/reqlite/internal/sqlite"
	"github.com/go-redis/redismock/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func TestClientGetNameOK(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	want := "username"
	mock.ExpectClientGetName().SetVal(want)
	function := getname.New(rdb)

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("client_getname", function); err != nil {
			return sqlite.SQLITE_ERROR, err
		}
		return sqlite.SQLITE_OK, nil
	})

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT CLIENT_GETNAME()")
	err = row.Err()
	if err != nil {
		t.Fatal(err)
	}

	var res string
	err = row.Scan(&res)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Compare(res, want) == 0 {
		t.Error(fmt.Sprintf("result : %s != desired : %s", res, want))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
