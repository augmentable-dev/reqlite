# Redis Commands and Corresponding SQLite Functions

| Redis Command                 | SQLite Function
|-------------------------------|---
| ACL LOAD                      | 🚧
| ACL SAVE                      | 🚧
| ACL LIST                      | 🚧
| ACL USERS                     | 🚧
| ACL GETUSER                   | 🚧
| ACL SETUSER                   | 🚧
| ACL DELUSER                   | 🚧
| ACL CAT                       | 🚧
| ACL GENPASS                   | 🚧
| ACL WHOAMI                    | 🚧
| ACL LOG                       | 🚧
| ACL HELP                      | 🚧
| APPEND                        | 🚧
| AUTH                          | 🚧
| BGREWRITEAOF                  | 🚧
| BGSAVE                        | 🚧
| BITCOUNT                      | ✅ [`BITCOUNT`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/bitcount)
| BITFIELD                      | 🚧
| BITOP                         | 🚧
| BITPOS                        | ✅ [`BITPOS`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/bitpos)
| BLPOP                         | 🚧
| BRPOP                         | 🚧
| BRPOPLPUSH                    | 🚧
| BLMOVE                        | 🚧
| BZPOPMIN                      | 🚧
| BZPOPMAX                      | 🚧
| CLIENT CACHING                | 🚧
| CLIENT ID                     | ✅ [`CLIENT_ID`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/client_id)
| CLIENT INFO                   | 🚧
| CLIENT KILL                   | 🚧
| CLIENT LIST                   | 🚧
| CLIENT GETNAME                | 🚧
| CLIENT GETREDIR               | 🚧
| CLIENT UNPAUSE                | 🚧
| CLIENT PAUSE                  | 🚧
| CLIENT REPLY                  | 🚧
| CLIENT SETNAME                | 🚧
| CLIENT TRACKING               | 🚧
| CLIENT TRACKINGINFO           | 🚧
| CLIENT UNBLOCK                | 🚧
| CLUSTER ADDSLOTS              | 🚧
| CLUSTER BUMPEPOCH             | 🚧
| CLUSTER COUNT-FAILURE-REPORTS | 🚧
| CLUSTER COUNTKEYSINSLOT       | 🚧
| CLUSTER DELSLOTS              | 🚧
| CLUSTER FAILOVER              | 🚧
| CLUSTER FLUSHSLOTS            | 🚧
| CLUSTER FORGET                | 🚧
| CLUSTER GETKEYSINSLOT         | 🚧
| CLUSTER INFO                  | 🚧
| CLUSTER KEYSLOT               | 🚧
| CLUSTER MEET                  | 🚧
| CLUSTER MYID                  | 🚧
| CLUSTER NODES                 | 🚧
| CLUSTER REPLICATE             | 🚧
| CLUSTER RESET                 | 🚧
| CLUSTER SAVECONFIG            | 🚧
| CLUSTER SET-CONFIG-EPOCH      | 🚧
| CLUSTER SETSLOT               | 🚧
| CLUSTER SLAVES                | 🚧
| CLUSTER REPLICAS              | 🚧
| CLUSTER SLOTS                 | 🚧
| COMMAND                       | 🚧
| COMMAND COUNT                 | 🚧
| COMMAND GETKEYS               | 🚧
| COMMAND INFO                  | 🚧
| CONFIG GET                    | ✅ [`CONFIG_GET`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/config_get)
| CONFIG REWRITE                | 🚧
| CONFIG SET                    | 🚧
| CONFIG RESETSTAT              | 🚧
| COPY                          | 🚧
| DBSIZE                        | ✅ [`DBSIZE`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/dbsize)
| DEBUG OBJECT                  | 🚧
| DEBUG SEGFAULT                | 🚧
| DECR                          | 🚧
| DECRBY                        | 🚧
| DEL                           | 🚧
| DISCARD                       | 🚧
| DUMP                          | 🚧
| ECHO                          | 🚧
| EVAL                          | 🚧
| EVALSHA                       | 🚧
| EXEC                          | 🚧
| EXISTS                        | 🚧
| EXPIRE                        | 🚧
| EXPIREAT                      | 🚧
| FAILOVER                      | 🚧
| FLUSHALL                      | 🚧
| FLUSHDB                       | 🚧
| GEOADD                        | 🚧
| GEOHASH                       | 🚧
| GEOPOS                        | 🚧
| GEODIST                       | 🚧
| GEORADIUS                     | 🚧
| GEORADIUSBYMEMBER             | 🚧
| GEOSEARCH                     | 🚧
| GEOSEARCHSTORE                | 🚧
| GET                           | 🚧
| GETBIT                        | 🚧
| GETDEL                        | 🚧
| GETEX                         | 🚧
| GETRANGE                      | 🚧
| GETSET                        | 🚧
| HDEL                          | 🚧
| HELLO                         | 🚧
| HEXISTS                       | 🚧
| HGET                          | 🚧
| HGETALL                       | ✅ [`HGETALL`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/hgetall)
| HINCRBY                       | 🚧
| HINCRBYFLOAT                  | 🚧
| HKEYS                         | 🚧
| HLEN                          | 🚧
| HMGET                         | 🚧
| HMSET                         | 🚧
| HSET                          | 🚧
| HSETNX                        | 🚧
| HRANDFIELD                    | 🚧
| HSTRLEN                       | 🚧
| HVALS                         | 🚧
| INCR                          | 🚧
| INCRBY                        | 🚧
| INCRBYFLOAT                   | 🚧
| INFO                          | 🚧
| LOLWUT                        | 🚧
| KEYS                          | 🚧
| LASTSAVE                      | 🚧
| LINDEX                        | 🚧
| LINSERT                       | 🚧
| LLEN                          | ✅ [`LLEN`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/llen)
| LPOP                          | 🚧
| LPOS                          | 🚧
| LPUSH                         | 🚧
| LPUSHX                        | 🚧
| LRANGE                        | ✅ [`LRANGE`](https://github.com/augmentable-dev/reqlite/tree/main/internal/redis/lrange)
| LREM                          | 🚧
| LSET                          | 🚧
| LTRIM                         | 🚧
| MEMORY DOCTOR                 | 🚧
| MEMORY HELP                   | 🚧
| MEMORY MALLOC-STATS           | 🚧
| MEMORY PURGE                  | 🚧
| MEMORY STATS                  | 🚧
| MEMORY USAGE                  | 🚧
| MGET                          | 🚧
| MIGRATE                       | 🚧
| MODULE LIST                   | 🚧
| MODULE LOAD                   | 🚧
| MODULE UNLOAD                 | 🚧
| MONITOR                       | 🚧
| MOVE                          | 🚧
| MSET                          | 🚧
| MSETNX                        | 🚧
| MULTI                         | 🚧
| OBJECT                        | 🚧
| PERSIST                       | 🚧
| PEXPIRE                       | 🚧
| PEXPIREAT                     | 🚧
| PFADD                         | 🚧
| PFCOUNT                       | 🚧
| PFMERGE                       | 🚧
| PING                          | 🚧
| PSETEX                        | 🚧
| PSUBSCRIBE                    | 🚧
| PUBSUB                        | 🚧
| PTTL                          | 🚧
| PUBLISH                       | 🚧
| PUNSUBSCRIBE                  | 🚧
| QUIT                          | 🚧
| RANDOMKEY                     | 🚧
| READONLY                      | 🚧
| READWRITE                     | 🚧
| RENAME                        | 🚧
| RENAMENX                      | 🚧
| RESET                         | 🚧
| RESTORE                       | 🚧
| ROLE                          | 🚧
| RPOP                          | 🚧
| RPOPLPUSH                     | 🚧
| LMOVE                         | 🚧
| RPUSH                         | 🚧
| RPUSHX                        | 🚧
| SADD                          | 🚧
| SAVE                          | 🚧
| SCARD                         | 🚧
| SCRIPT DEBUG                  | 🚧
| SCRIPT EXISTS                 | 🚧
| SCRIPT FLUSH                  | 🚧
| SCRIPT KILL                   | 🚧
| SCRIPT LOAD                   | 🚧
| SDIFF                         | 🚧
| SDIFFSTORE                    | 🚧
| SELECT                        | 🚧
| SET                           | 🚧
| SETBIT                        | 🚧
| SETEX                         | 🚧
| SETNX                         | 🚧
| SETRANGE                      | 🚧
| SHUTDOWN                      | 🚧
| SINTER                        | 🚧
| SINTERSTORE                   | 🚧
| SISMEMBER                     | 🚧
| SMISMEMBER                    | 🚧
| SLAVEOF                       | 🚧
| REPLICAOF                     | 🚧
| SLOWLOG                       | 🚧
| SMEMBERS                      | 🚧
| SMOVE                         | 🚧
| SORT                          | 🚧
| SPOP                          | 🚧
| SRANDMEMBER                   | 🚧
| SREM                          | 🚧
| STRALGO                       | 🚧
| STRLEN                        | 🚧
| SUBSCRIBE                     | 🚧
| SUNION                        | 🚧
| SUNIONSTORE                   | 🚧
| SWAPDB                        | 🚧
| SYNC                          | 🚧
| PSYNC                         | 🚧
| TIME                          | 🚧
| TOUCH                         | 🚧
| TTL                           | 🚧
| TYPE                          | 🚧
| UNSUBSCRIBE                   | 🚧
| UNLINK                        | 🚧
| UNWATCH                       | 🚧
| WAIT                          | 🚧
| WATCH                         | 🚧
| ZADD                          | 🚧
| ZCARD                         | 🚧
| ZCOUNT                        | 🚧
| ZDIFF                         | 🚧
| ZDIFFSTORE                    | 🚧
| ZINCRBY                       | 🚧
| ZINTER                        | 🚧
| ZINTERSTORE                   | 🚧
| ZLEXCOUNT                     | 🚧
| ZPOPMAX                       | 🚧
| ZPOPMIN                       | 🚧
| ZRANDMEMBER                   | 🚧
| ZRANGESTORE                   | 🚧
| ZRANGE                        | 🚧
| ZRANGEBYLEX                   | 🚧
| ZREVRANGEBYLEX                | 🚧
| ZRANGEBYSCORE                 | 🚧
| ZRANK                         | 🚧
| ZREM                          | 🚧
| ZREMRANGEBYLEX                | 🚧
| ZREMRANGEBYRANK               | 🚧
| ZREMRANGEBYSCORE              | 🚧
| ZREVRANGE                     | 🚧
| ZREVRANGEBYSCORE              | 🚧
| ZREVRANK                      | 🚧
| ZSCORE                        | 🚧
| ZUNION                        | 🚧
| ZMSCORE                       | 🚧
| ZUNIONSTORE                   | 🚧
| SCAN                          | 🚧
| SSCAN                         | 🚧
| HSCAN                         | 🚧
| ZSCAN                         | 🚧
| XINFO                         | 🚧
| XADD                          | 🚧
| XTRIM                         | 🚧
| XDEL                          | 🚧
| XRANGE                        | 🚧
| XREVRANGE                     | 🚧
| XLEN                          | 🚧
| XREAD                         | 🚧
| XGROUP                        | 🚧
| XREADGROUP                    | 🚧
| XACK                          | 🚧
| XCLAIM                        | 🚧
| XAUTOCLAIM                    | 🚧
| XPENDING                      | 🚧
| LATENCY DOCTOR                | 🚧
| LATENCY GRAPH                 | 🚧
| LATENCY HISTORY               | 🚧
| LATENCY LATEST                | 🚧
| LATENCY RESET                 | 🚧
| LATENCY HELP                  | 🚧