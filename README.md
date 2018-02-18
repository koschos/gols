# Overview

Simple link short REST API

## Install

Clone repository
```bash
git clone git@github.com:koschos/gols.git
```

Install binary
```bash
go install github.com/koschos/gols
```

## Configuration

You can configure database connection and HTTP port to listen.
Please use config.yaml file from example folder.
You can not change config file name or format.

```yaml
port: HTTP_PORT
db:
  host: YOUR_HOST
  port: YOUR_PORT
  user: YOUR_USER
  pass: YOUR_PASS
  name: YOUR_DB_NAME
```

## Hasher

Hash function for generating url hash is md5.

## Slug generator

Slug generated randomly with length 6 using base 36 (a-zA-Z0-9)

## Run

```bash
./gols
```

## Run unit tests

```bash
export DB_USERNAME=root
export DB_PASSWORD=my-secret
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_TEST_NAME=gols

go test ./...
```

## Test manually

New link created, code 201 Created

```bash
curl -i -X POST http://localhost:8080/api/v1/short-link/ -d '{ "url": "http://test.com" }'

HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Mon, 12 Feb 2018 22:40:01 GMT
Content-Length: 79

{"data":{"slug":"slug","url":"http://test.com","url_hash":"hash"},"status":201}
```

Link already exists, code 208 Already Reported

```bash
curl -i -X POST http://localhost:8080/api/v1/short-link/ -d '{ "url": "http://test.com" }'

HTTP/1.1 208 Already Reported
Content-Type: application/json; charset=utf-8
Date: Mon, 12 Feb 2018 22:40:01 GMT
Content-Length: 79

{"data":{"slug":"slug","url":"http://test.com","url_hash":"hash"},"status":201}
```