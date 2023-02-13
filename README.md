# Scheduling API example with golang

This is a repository with scheduling examples of HTTP requests with GIN and mySQL

## Requirements

**Go 1.16 >**
```
$ Get installer in: https://go.dev/learn/
```

**MySQL - MariaDB Database**
```
$ Get installers in:
- XAMPP: https://www.apachefriends.org/pt_br/index.html
- HeidiSQL: https://www.heidisql.com/download.php
```

## Getting Started



All environment variables are in `env.yaml`, configure them according to your database.

All executables are defined inside `main.go`.

```
$ go run main.go
```

The server should start at: [`http:localhost:8080`](http:localhost:8080)

## Glossary

1. Database: SQL File to create your db.
2. Domain: Entities, model and usecase
3. Infrastructure: Basic configuration, Database connection, handler and repository
4. Interfaces: All project interfaces
5. Utils: Packages that are used throughout the project

## Architecture Goals

1. Follow the concepts of the clean architecture, with extensive use of interfaces.


## Available Endpoint

### `POST /agendas`

- Create a  new schedule.

```
curl --location --request POST 'http://localhost:8080/agendas' \
--header 'Content-Type: application/json' \
--data-raw '{
    "empresa": {
        "cnpj": "26488705000193"
    },
    "horario": "09:00"
}'
```

### `GET /agendas`

- Get all schedules

```
curl --location --request GET 'http://localhost:8080/agendas' \
--header 'Content-Type: application/json'
```


### `GET /agendas:disponibilidade`

- Get all schedules and availability


```
curl --location --request GET 'http://localhost:8080/agendas:disponibilidade' \
--header 'Content-Type: application/json'
```

