#!/bin/bash

export GOPATH=${GOPATH}:`pwd`
cd src/test_pkg
go test -v