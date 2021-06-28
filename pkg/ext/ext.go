package ext

import (
	"os"

	"github.com/augmentable-dev/reqlite/internal/json/get"
	"github.com/augmentable-dev/reqlite/internal/json/mget"
	"github.com/augmentable-dev/reqlite/internal/redis/bitcount"
	"github.com/augmentable-dev/reqlite/internal/redis/bitpos"
	"github.com/augmentable-dev/reqlite/internal/redis/client_id"
	"github.com/augmentable-dev/reqlite/internal/redis/config_get"
	"github.com/augmentable-dev/reqlite/internal/redis/dbsize"
	"github.com/augmentable-dev/reqlite/internal/redis/dump"
	"github.com/augmentable-dev/reqlite/internal/redis/hgetall"
	"github.com/augmentable-dev/reqlite/internal/redis/llen"
	"github.com/augmentable-dev/reqlite/internal/redis/lrange"
	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
	"go.riyazali.net/sqlite"
)

func init() {
	options := &redis.Options{}
	if net := os.Getenv("REQLITE_NET"); net != "" {
		options.Network = net
	}
	if addr := os.Getenv("REQLITE_ADDR"); addr != "" {
		options.Addr = addr
	}
	if user := os.Getenv("REQLITE_USER"); user != "" {
		options.Username = user
	}
	if pass := os.Getenv("REQLITE_PASS"); pass != "" {
		options.Password = pass
	}
	rdb := redis.NewClient(options)
	// TODO how should closing the client be handled?
	// is it necessary?

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateModule("lrange", lrange.New(rdb),
			sqlite.EponymousOnly(true), sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateModule("hgetall", hgetall.New(rdb),
			sqlite.EponymousOnly(true), sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("bitcount", bitcount.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("bitpos", bitpos.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("client_id", client_id.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateModule("config_get", config_get.New(rdb),
			sqlite.EponymousOnly(true), sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("dbsize", dbsize.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("dump", dump.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("llen", llen.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("json_get", get.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateModule("json_mget", mget.New(rdb)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}
