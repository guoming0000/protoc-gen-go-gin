//go:generate protoc -I. -I ./third_party --go-gin_out=./ --go_out=./ api/article/article.proto
//go:generate protoc --go-errors_out=./ api/article/article_ecode.proto

package main
