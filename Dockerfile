FROM golang:1.19.3-alpine3.16 AS build
WORKDIR /src
COPY . /src
ARG VERSION
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -ldflags="-s -w -X 'main.version=$VERSION'" \
    -o sntl-gw

FROM alpine:3.16 AS src
COPY --from=build /src/sntl-gw .
ENTRYPOINT /sntl-gw
