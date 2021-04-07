.PHONY: build clean deploy

run:
	go run cmd/server/main.go

build:
	go mod tidy
	env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/server/main.go

test:
	go test ./... -v

docs:
	cd cmd/server/;swag init

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose --aws-s3-accelerate
