#!/bin/bash
#
# lxy@20230612
#
set -ve
cd cmd/protoc-gen-openapiv2
go build
cd ../../

#protoc --proto_path=. \
#        --proto_path=./third_party \
#        --openapiv2_out . \
#        --openapiv2_opt logtostderr=true \
#        --openapiv2_opt json_names_for_fields=false \
#        api/article/article.proto

protoc --proto_path=. \
        --proto_path=./third_party \
        --openapiv2_out . \
        --openapiv2_opt logtostderr=true \
        --openapiv2_opt json_names_for_fields=false \
        api/article/auth.proto