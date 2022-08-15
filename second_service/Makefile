CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

gen-proto-module:
	./scripts/gen_proto.sh ${CURRENT_DIR}

run:
	go run cmd/main.go
