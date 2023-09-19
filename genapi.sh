#!/bin/bash
#
# lxy@20230612
#
set -ve
go build

protoc -I. -I ./third_party --go-gin_out=./ --go_out=./ api/article.proto
protoc -I. -I ./third_party --go-gin_out=./ --go_out=./ api/auth.proto
protoc --go-errors_out=./ api/article_ecode.proto

