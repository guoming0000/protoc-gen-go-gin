//go:generate protoc -I. -I ./third_party --go-gin_out=./ --go_out=./ --validate_out=lang=go:./ api/article.proto
//go:generate protoc --go-errors_out=./ api/article_error.proto

package main
