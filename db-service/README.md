# SQL Sandbox

1. `cd db-service`
2. `make db`

Connect to container: `docker exec -it db-service-database-1 psql -U user db`

## PostgreSQL & Go

### Drivers 
1. github.com/lib/pq (old) – Pure Go Postgres driver for database/sql
2. github.com/jackc/pgx (new) - PostgreSQL driver and toolkit for Go (new). 
   * Better perfomance, logging, no panics.

### Interfaces
1. [database/sql](http://go-database-sql.org/overview.html) (old) - Package sql provides a generic interface around SQL (or SQL-like) databases. Transactions, queries, etc. 
2. [jmoiron/sqlx](http://jmoiron.github.io/sqlx/) (new) – extension on database/sql. 
   * Allows to use structs on top of query instead of marshalling each column to Go variable.
3. [sqlc](https://docs.sqlc.dev/) - sqlc generates fully type-safe idiomatic Go code from SQL
   

### Bouncer
1. [pgbouncer](https://www.pgbouncer.org/features.html) - creates a pool of database connections and provides these connections to clients when the connection is required. Thus, if you have a lot of client connections to the database, PgBouncer can reduce the number of PostgreSQL backends processes.


## Logging

1. https://golang.org/pkg/database/sql/#DB.Stats
2. https://godoc.org/github.com/jackc/pgx#ConnPool.Stat
