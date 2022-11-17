# sentential-gw

[![main](https://github.com/wheegee/sentential-gw/actions/workflows/main.yml/badge.svg)](https://github.com/wheegee/sentential-gw/actions/workflows/main.yml)

Local HTTP => Lambda invocation proxy designed to replicate the behavior of a Lambda public URL for local verification and testing.

## Development

### Gateway

```sh
go mod tidy
go run main.go
```

### Sentential

In a different directory, create a [sntl](https://github.com/wheegee/sentential) project for testing

Deploy said `sntl` project locally with:
```sh
sntl deploy local
```

## Release

1. Merge pull request
1. Push a [semantically versioned](https://semver.org) tag to `main`
