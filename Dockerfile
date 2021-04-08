### Bulder
FROM golang:1.16.3-alpine3.13 as builder

RUN apk update; \
    apk add git ca-certificates upx

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod tidy
# install dependencies

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o bin/main cmd/server/main.go; \
    upx --best --lzma bin/main
# compile & pack

### Executable Image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/src/app/bin/main ./main

EXPOSE 8080

ENTRYPOINT ["./main"]