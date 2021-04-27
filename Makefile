#!/bin/bash

all:fmt test

# 运行所有的测试用例
test:
	go test -v ./*.go -cover
	go test -v ./... -cover

# 安装 goconvey 框架
install:
	go get github.com/smartystreets/goconvey

# 运行 convey web UI
convey:install
	goconvey -port 2021

fmt:
	go fmt ./*.go
