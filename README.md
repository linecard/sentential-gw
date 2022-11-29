# sentential-gw

[![Main](https://github.com/wheegee/sentential-gw/actions/workflows/main.yml/badge.svg)](https://github.com/wheegee/sentential-gw/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/wheegee/sentential-gw)](https://goreportcard.com/report/github.com/wheegee/sentential-gw)

Local HTTP => Lambda invocation proxy designed to replicate the behavior of a Lambda public URL for local verification and testing.

## Usage

Intended to be utilized alongside [sentential](https://github.com/wheegee/sentential).

## Development

### Gateway

```sh
go mod tidy
go run main.go
```

### Sentential

In a different directory, create a [sntl](https://github.com/wheegee/sentential) project for testing.

Deploy said `sntl` project locally with:
```sh
sntl deploy local
```
