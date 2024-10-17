MAIN:=./cmd/server
TARGET:=base-code
WIN_TARGET:=base-code.exe
SERVER_CMD:=./${TARGET}
PROTOC_DEST:=./
PROTOC_FLAGS:=--go_out=${PROTOC_DEST} --go_opt=paths=source_relative --go-grpc_out=${PROTOC_DEST} --go-grpc_opt=paths=source_relative
USER_PROTO_FILES:=./grpc/user/user.proto

# example migration create command -
# migrate create -ext sql -seq -dir migrations create-some-table

build-proto:
	protoc ${PROTOC_FLAGS} ${USER_PROTO_FILES}

run-server:
	${SERVER_CMD}

tidy:
	go mod tidy

install-proto-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

install-dev-deps:
	go install github.com/air-verse/air@latest

install-deps:
	go mod download

prepare: install-proto-deps install-dev-deps install-deps tidy

dev: prepare
	air

build: install-deps
	go build -o ${TARGET} ${MAIN}

start: build run-server

dev-win: prepare
	air -c .air.win.toml


build-win: install-deps
	go build -o ${WIN_TARGET} ${MAIN}


start-win: build-win run-server