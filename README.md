# Hotel API Exercise with golang

This is a repository with a REST API made with GIN and PostgresSQL

## Requirements

**Go 1.16 >**
```
$ Get installer in: https://go.dev/learn/
```

**PostgreSQL**
```
$ Get installers in:
- postgreSQL: https://www.postgresql.org/
- pgAdmin: https://www.pgadmin.org/ (optional)
- dbeaver: https://dbeaver.io/download/ (optional)
```

## Getting Started

All environment variables are in `env.yaml`, configure them according to your database.

All executables are defined inside `cmd/main.go`.

Run from base folder to envs be read correctly.
```
$ go run cmd/main.go
```

The server should start at: [`http:localhost:8080`](http:localhost:8080)

## Glossary

1. Database: SQL File to create your db.
2. Domain: Entities, model and usecase
3. Infrastructure: Basic configuration, builds, database connection, handler and repository
4. Interfaces: All project interfaces
5. Utils: Packages that are used throughout the project

## Architecture Goals

1. Follow the concepts of the clean architecture, with extensive use of interfaces.