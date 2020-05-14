
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	@protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/user/user.proto
    

.PHONY: build
build: proto
	@echo building user-srv...
	@go build -o ../../build/user-srv *.go
	@echo building success

.PHONY: test
test:
	go test -v ./... -cover

run:
	@../../build/user-srv --registry=consul

.PHONY: docker
docker:
	docker build . -t user-service:latest
