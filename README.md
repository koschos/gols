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

TBD

database name must be gols

## Hasher

TBD

## Slug generator

TBD

## Run

```bash
./gols
```

## Run unit tests

```bash
go test
```

## Test manually

```bash
curl -i -X POST http://localhost:8080/api/v1/short-link/ -d '{ "url": "http://test.com" }'

HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Mon, 12 Feb 2018 22:40:01 GMT
Content-Length: 79

{"data":{"slug":"slug","url":"http://test.com","url_hash":"hash"},"status":201}
```