# reqlite

`reqlite` makes it possible to query data in redis with SQL.
Queries are excuted client-side with SQLite (not on the redis server).
This project is distributed as a SQLite [run-time loadable extension](https://www.sqlite.org/loadext.html) and a standalone binary.

## LRANGE

```sql
SELECT * FROM LRANGE('some-key', 0, 10)
```
