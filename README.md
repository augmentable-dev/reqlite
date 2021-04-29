[![BuildStatus](https://github.com/augmentable-dev/reqlite/workflows/tests/badge.svg)](https://github.com/augmentable-dev/reqlite/actions?workflow=tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/augmentable-dev/reqlite)](https://goreportcard.com/report/github.com/augmentable-dev/reqlite)
# reqlite

`reqlite` makes it possible to query data in redis with SQL.
Queries are executed client-side with SQLite (not on the redis server).
This project is distributed as a SQLite [run-time loadable extension](https://www.sqlite.org/loadext.html) and a standalone binary.

## Getting Started

To build a run-time loadable extension, run `make` in the root of the source tree.
The `reqlite.so` file should be in `.build/reqlite.so`, which you can use immediately in a [SQLite shell](https://sqlite.org/cli.html):

```sql
sqlite3
sqlite> .load .build/reqlite.so
sqlite> SELECT * FROM LRANGE('some-key', 0, 10);
```

## Table-Valued Functions

The following [table-valued functions](https://www.sqlite.org/vtab.html#tabfunc2) are based on corresponding Redis commands.

### LRANGE

```sql
SELECT * FROM LRANGE('some-key', 0, 10)
```

### HGETALL

```sql
SELECT * FROM HGETALL('myhash')
```
