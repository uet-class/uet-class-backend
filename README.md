# UET Class: Backend

*Backend repository for the UET Class project*.

## Pre-requisite

- `Go` version: `>= 1.19`

## Usage

__NOTE: These steps below are instructions for Linux environment. If you are using another type of OSes (e.g Windows, MacOS, ...), please just use these steps as references.__

- Clone the repository and `cd` inside:
  ``` bash
  git clone https://github.com/UET-Class/uet-class-backend.git
  cd uet-class-backend
  ```

- Install dependencies:
  ```bash
  go install
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

| Variables     | Usage                           |
| ------------- | ------------------------------- |
| `SERVER_HOST` | Server hostname                 |
| `SERVER_PORT` | Server port number              |
| `DB_HOST`     | Database hostname               |
| `DB_PORT`     | Database port number            |
| `DB_USER`     | Database user to connect        |
| `DB_PASSWORD` | Password of the user to connect |
| `DB_NAME`     | Database name                   |

These variables should be defined in a YAML file, named `develop.yaml`, and placed in the `config` directory.
For the real configuration, please refer to [configuration-management](https://github.com/uet-class/configuration-management) repository.
Example: 

``` yaml
# config/develop.yaml
SERVER_HOST: localhost
SERVER_PORT: :8080  # Note that there is a colon ':' before the port number
DB_HOST: localhost
DB_PORT: 5432
DB_USER: user
DB_PASSWORD: passwd
DB_NAME: db_name
```

## Seeding data

The seeding script is placed in `seed/` directory of this repo.

``` bash
go run seed/seed.go
```
## To-do
