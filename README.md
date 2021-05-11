[![BuildStatus](https://github.com/augmentable-dev/reqlite/workflows/tests/badge.svg)](https://github.com/augmentable-dev/reqlite/actions?workflow=tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/augmentable-dev/reqlite)](https://goreportcard.com/report/github.com/augmentable-dev/reqlite)
# reqlite

`reqlite` makes it possible to query data in [Redis](https://redis.io/) with [SQL](https://sqlite.org/lang.html).
Queries are executed client-side with SQLite (not on the redis server).
This project is distributed as a SQLite [run-time loadable extension](https://www.sqlite.org/loadext.html) and (soon) as a standalone binary (CLI).

This project is experimental for the time being.
It's made possible by a [great library for building SQLite extensions in go](https://github.com/riyaz-ali/sqlite).

The [JSON1 extension](https://www.sqlite.org/json1.html) is also included by default as a convenience.

<p align="center">
    <img src="./diagram.png?raw=true">
</p>

## Use Cases

What can or should I use this for?
This project is pretty experimental and part of that is exploring use-cases to understand what's possible and interesting!

A common situation is a [task queue](https://redislabs.com/ebook/part-2-core-concepts/chapter-6-application-components-in-redis/6-4-task-queues/) in Redis.
If you're [using a `LIST`](https://redislabs.com/ebook/part-2-core-concepts/chapter-6-application-components-in-redis/6-4-task-queues/6-4-1-first-in-first-out-queues/) as a queue holding JSON objects, `reqlite` + the SQLite json1 extension could be used to issue basic "slicing and dicing" queries against your task queue.

```sql
-- what are the most common tasks currently in the queue?
SELECT count(*), json_extract(value, '$.task') as task
FROM LRANGE('my-queue', 0, 100)
GROUP BY task
ORDER BY count(*) DESC
```

In general, Redis is fairly accessible from [many programming languages](https://redis.io/clients), and any query using `reqlite` could probably be implemented in a language of your choice using a Redis client.
However, sometimes declarative SQL can be a better choice to express what you're looking for, and that's where this project may be most useful.
Since `reqlite` is distributed as a run-time loadable SQLite extension, it can be loaded into a language using a SQLite driver as well, which would allow you to mix SQL and the "host" language to access data in Redis.

## Getting Started

To build a run-time loadable extension, run `make` in the root of the source tree.
The `reqlite.so` file should be in `.build/reqlite.so`, which you can use immediately in a [SQLite shell](https://sqlite.org/cli.html):

```sql
sqlite3
sqlite> .load .build/reqlite.so
sqlite> SELECT * FROM LRANGE('some-key', 0, 10);
```

### Connecting to Redis

Currently, the Redis connection can only be set via the following `env` variables:

| ENV          | Default          | Description                           |
|--------------|------------------|---------------------------------------|
| REQLITE_NET  | `tcp`            | Network type - either `tcp` or `udp`  |
| REQLITE_ADDR | `localhost:6379` | Network address of the redis instance |
| REQLITE_USER | (none)           | Redis username                        |
| REQLITE_PASS | (none)           | Redis password                        |

TODO - Implement another mechanism (SQLite UDFs?) for setting up the connection information.

## Commands

Currently, only read operations are targeted to be implemented as SQLite [scalar functions](https://www.sqlite.org/appfunc.html) or [table-valued functions](https://www.sqlite.org/vtab.html#tabfunc2).
In the examples below, you'll see how a SQLite scalar or table-valued function maps to a corresponding [Redis command](https://redis.io/commands), based on the response type.
Note that there won't always be an exact correspondence, and currently not all Redis commands are targeted to be implemented (read-only for now).

```sql
SELECT * FROM some_table_valued_function('param', 1, 2) -- function that returns a table
SELECT some_scalar_function('param', 1, 2) -- function that returns a scalar value
```

Available functions are listed below.
For a full list of Redis commands and corresponding SQLite functions, [see here](https://github.com/augmentable-dev/reqlite/tree/main/commands.md).

### LRANGE

```sql
SELECT * FROM LRANGE('some-key', 0, 10)
```

### HGETALL

```sql
SELECT * FROM HGETALL('myhash')
```

### BITCOUNT

```sql
SELECT BITCOUNT('some-key')
SELECT BITCOUNT('some-key', 1, 1)
```

### BITPOS

```sql
SELECT BITPOS('some-key', 0)
SELECT BITPOS('some-key', 1, 2)
```

### CLIENT ID

```sql
SELECT CLIENT_ID()
```

### CONFIG GET

```sql
SELECT * FROM CONFIG_GET('*max-*-entries*')
SELECT * FROM CONFIG_GET -- equivalent to CONFIG GET *
```

### RedisJSON ([link](https://oss.redislabs.com/redisjson/))

#### JSON_GET

```sql
SELECT JSON_GET('my-json-key')
SELECT JSON_GET('my-json-key', 'some.path[2]')
```

#### JSON_MGET

```sql
SELECT * FROM JSON_MGET('some.path', 'key1,key2,key3')
```
