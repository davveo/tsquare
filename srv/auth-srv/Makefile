
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/auth/auth.proto
    

.PHONY: build
build: proto
	@echo building auth-srv...
	@go build -o ../../build/auth-srv *.go
	@echo building success

.PHONY: test
test:
	go test -v ./... -cover
run:
	@../../build/auth-srv --registry=consul

.PHONY: docker
docker:
	docker build . -t auth-service:latest
