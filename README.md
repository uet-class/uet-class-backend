# UET Class: Backend

*Backend repository for the UET Class project*.

## Pre-requisite

- `Go` version: `>= 1.19`

## Usage

- Clone the repository and `cd` inside:
  ``` bash
  git clone https://github.com/UET-Class/uet-class-backend.git
  cd uet-class-backend
  ```

- Build the source code:
  ``` bash
  go build main.go
  ```

- Run the binary compiled file:
  ``` bash
  ./main
  ```

  This will start the server at `http://localhost:8080`.

## Environment variables

| Variables           | Usage                           |
| ------------------- | ------------------------------- |
| `SERVER_HOST`       | Server hostname                 |
| `SERVER_PORT`       | Server port number              |
| `POSTGRES_HOST`     | Database hostname               |
| `POSTGRES_PORT`     | Database port number            |
| `POSTGRES_USER`     | Database user to connect        |
| `POSTGRES_PASSWORD` | Password of the user to connect |
| `POSTGRES_DATABASE` | Postgres name                   |
| `REDIS_HOST`        | Redis hostname                  |
| `REDIS_PORT`        | Redis port number               |
| `REDIS_PASSWORD`    | Redis password (if any)         |
| `REDIS_DATABASE`    | Redis database name             |


These variables should be defined in a YAML file, named `develop.yaml`, and placed in the `config` directory of this repository.
For the real configuration, please refer to [configuration-management](https://github.com/uet-class/configuration-management) repository.
Example: 

``` yaml
# config/develop.yaml
---
SERVER_HOST: localhost
SERVER_PORT: :8080

POSTGRES_HOST: localhost
POSTGRES_PORT: 15432
POSTGRES_USER: uc_root
POSTGRES_PASSWORD: uc_pwd
POSTGRES_DATABASE: uet_class_dev

REDIS_HOST: localhost
REDIS_PORT: 6379
REDIS_PASSWORD: ""
REDIS_DATABASE: 0
```

## Seeding data

The seeding script is placed in `seed/` directory of this repo.

``` bash
go run seed/seed.go
```
## To-do
