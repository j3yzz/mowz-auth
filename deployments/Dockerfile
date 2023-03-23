FROM golang:alpine

WORKDIR /usr/src/app


COPY . .

RUN go install github.com/cosmtrek/air@latest

WORKDIR /tmp/go-migrate

ARG VERSION="4.15.2"

RUN set -x \
    && apk add --no-cache git \
    && git clone --branch "v${VERSION}" --depth 1 --single-branch https://github.com/golang-migrate/migrate /tmp/go-migrate

RUN set -x \
    && CGO_ENABLED=0 go build -tags 'mysql' -ldflags="-s -w" -o ./migrate ./cmd/migrate \
    && ./migrate -version

RUN go mod tidy

RUN cp /tmp/go-migrate/migrate /usr/bin/migrate

WORKDIR /usr/src/app

CMD ["air", "-c", ".air.toml"]
