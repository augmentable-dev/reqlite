package ext

import (
	"os"

	"github.com/augmentable-dev/reqlite/internal/hgetall"
	"github.com/augmentable-dev/reqlite/internal/lrange"
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

	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateModule("lrange", lrange.New(rdb),
			sqlite.EponymousOnly(true), sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateModule("hgetall", hgetall.New(rdb),
			sqlite.EponymousOnly(true), sqlite.ReadOnly(true)); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}
