//go:generate protoc -I. -I ./third_party --go-gin_out=./ --go_out=./ --validate_out=lang=go:./ api/article.proto

package main
