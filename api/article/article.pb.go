// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/article.proto

package article

import (
	_ "github.com/golang/protobuf/ptypes/any"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// 消息名使用首字母大写驼峰风格(CamelCase)，例如message StudentReq { ... }
type GetArticlesReq struct {
	// 文章标题
	Title string `json:"title" binding:"required"` // 标题名称右注释
	Page  int32  `json:"page,omitempty" binding:"lt=999"`
	// 字段名使用小写下划线的风格，例如 string status_code = 1
	PageSize int32 `json:"page_size,omitempty" binding:"lte=101"`
	// 作者id
	AuthorId int32                    `json:"author_id,omitempty"`
	Email    string                   `json:"email,omitempty" binding:"email"`
	Name     string                   `json:"name,omitempty"`
	Home     *GetArticlesReq_Location `json:"home" binding:"required"`
	TestStr  string                   `json:"test_str" binding:"required"`
}

// 重要注释1
type GetArticlesReply struct {
	Total    int64      `json:"total" binding:"required"`
	Articles []*Article `json:"articles"`
}

type Article struct {
	// 标题名称上注释
	Title    string `json:"title,omitempty"` // 标题名称右注释
	Content  string `json:"content" binding:"required"`
	AuthorId int32  `json:"author_id,omitempty"`
}

type GetArticlesReq_Location struct {
	Lat float64 `json:"lat,omitempty" binding:"lte:90,gte:-90"`
	Lng float64 `json:"lng,omitempty" binding:"gte:-180,lte:180"`
}
