# UET Class: Backend

*Backend repository for the UET Class project*.

## Pre-requisite

- `Go` version: `>= 1.19`
- `Make` (optional, if you want to make use of the Makefile)

## Usage

- Clone the repository and `cd` inside:
  ``` bash
  git clone https://github.com/UET-Class/uet-class-backend.git
  cd uet-class-backend
  ```

- Start server at `http://localhost:8080`:
  ``` bash
  go run main.go
  ```

  OR, you can use `make`:
    ``` make
    make
    ```

## Environment variables

``` ini
# .env
UC_DOMAIN_NAME=uetclass-dev.duckdns.org
GCP_PROJECT_ID=uet-class
GCS_BUCKET_LOCATION=asia
GCS_BUCKET_CLASS_PREFIX=uc-class
GCS_BUCKET_AVATARS=uc-avatars
GOOGLE_APPLICATION_CREDENTIALS=storage/gcs-sa.json

SERVER_HOST=localhost
SERVER_PORT==8080

POSTGRES_HOST=localhost
POSTGRES_PORT=15432
POSTGRES_USER=uc_root
POSTGRES_PASSWORD=xxx
POSTGRES_DATABASE=uet_class_dev

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DATABASE=0

SMTP_EMAIL_USERNAME=xxx@gmail.com
SMTP_EMAIL_PASSWORD=xxx
SMTP_HOSTNAME=smtp.gmail.com
SMTP_PORT=587
```

## Seeding data

The seeding script is placed in `seed/` directory of this repo.

``` bash
go run seed/seed.go
```
## To-do
