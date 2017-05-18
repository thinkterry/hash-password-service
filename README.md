# Hash password service

_HTTP service that hashes passwords using SHA512 and Base64 encoding_

To begin, start the server:

```bash
$ go run hash_password.go
```

Requests may then be made to http://localhost:8080. One of two requests is permitted:

1. Hash password
2. Shut down server

A delay of five seconds is applied to each response.

## Hash password

Hash a password via an HTTP POST to /, e.g.

```bash
$ curl -s --data "password=angryMonkey" http:/localhost:8080/
ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==
```

## Shut down server

Shut down the server via an HTTP POST to /shutdown, i.e.

```bash
$ curl -s -X POST http:/localhost:8080/shutdown
```
