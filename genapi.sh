#!/bin/bash
#
# lxy@20230612
#
set -ve
go build

protoc -I. -I ./third_party --go-gin_out=./ --go_out=./ api/article/article.proto
protoc -I. -I ./third_party --go-gin_out=./ --go_out=./ api/article/auth.proto
protoc --go-errors_out=./ api/article/article_ecode.proto

