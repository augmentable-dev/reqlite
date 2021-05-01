[![BuildStatus](https://github.com/augmentable-dev/reqlite/workflows/tests/badge.svg)](https://github.com/augmentable-dev/reqlite/actions?workflow=tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/augmentable-dev/reqlite)](https://goreportcard.com/report/github.com/augmentable-dev/reqlite)
# reqlite

`reqlite` makes it possible to query data in redis with SQL.
Queries are executed client-side with SQLite (not on the redis server).
This project is distributed as a SQLite [run-time loadable extension](https://www.sqlite.org/loadext.html) and (soon) as a standalone binary (CLI).

This project is experimental for the time being.
It's made possible by a [great library for building SQLite extensions in go](https://github.com/riyaz-ali/sqlite).

The [JSON1 extension](https://www.sqlite.org/json1.html) is also included by default as a convenience.

## Use Cases

What can or should I use this for?
This project is pretty experimental and part of that is exploring use-cases to understand what's possible or interesting!

A common situation is a [task queue](https://redislabs.com/ebook/part-2-core-concepts/chapter-6-application-components-in-redis/6-4-task-queues/) in Redis.
If you're [using a `LIST`](https://redislabs.com/ebook/part-2-core-concepts/chapter-6-application-components-in-redis/6-4-task-queues/6-4-1-first-in-first-out-queues/) as a queue holding JSON objects, `reqlite` + the SQLite json1 extension could be used to issue basic "slicing and dicing" queries against your task queue.

```sql
-- what are the most common tasks currently in the queue?
SELECT count(*), json_extract(value, '$.task') as task
FROM LRANGE('my-queue', 0, 100)
GROUP BY task
ORDER BY count(*) DESC
```

## Getting Started

To build a run-time loadable extension, run `make` in the root of the source tree.
The `reqlite.so` file should be in `.build/reqlite.so`, which you can use immediately in a [SQLite shell](https://sqlite.org/cli.html):

```sql
sqlite3
sqlite> .load .build/reqlite.so
sqlite> SELECT * FROM LRANGE('some-key', 0, 10);
```

## Commands

Currently, only read operations are targeted to be implemented as SQLite [scalar functions](https://www.sqlite.org/appfunc.html) or [table-valued functions](https://www.sqlite.org/vtab.html#tabfunc2).
In the examples below, you'll see how a SQLite scalar or table-valued function maps to a corresponding [Redis command](https://redis.io/commands), based on the response type.
Note that there won't always be an exact correspondence, and currently not all Redis commands are targeted to be implemented (read-only for now).

```sql
SELECT * FROM some_table_valued_function('param', 1, 2) -- function that returns a table
SELECT some_scalar_function('param', 1, 2) -- function that returns a scalar value
```

### LRANGE

```sql
SELECT * FROM LRANGE('some-key', 0, 10)
```

### HGETALL

```sql
SELECT * FROM HGETALL('myhash')
```

### RedisJSON ([link](https://oss.redislabs.com/redisjson/))

#### JSON_GET

```sql
SELECT JSON_GET('my-json-key')
```

```sql
SELECT JSON_GET('my-json-key', 'some.path[2]')
```

#### JSON_MGET

```sql
SELECT * FROM JSON_MGET('some.path', 'key1,key2,key3')
```
