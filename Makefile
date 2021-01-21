#!/bin/bash

test:
	go test -v ./...

install:
	go get github.com/smartystreets/goconvey

convey:install
	goconvey -port 2021
