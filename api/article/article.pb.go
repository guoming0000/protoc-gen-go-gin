// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/article/article.proto

package article

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// 消息名使用首字母大写驼峰风格(CamelCase)，例如message StudentReq { ... }
type GetArticlesReq struct {
	// 文章标题
	Title string `json:"title,omitempty" binding:"required"` // 标题名称右注释
	Page  int32  `json:"page,omitempty" binding:"lt=999"`
	// 字段名使用小写下划线的风格，例如 string status_code = 1
	PageSize int32 `json:"page_size,omitempty" binding:"lte=101"`
	// 作者id
	AuthorId int32                    `json:"author_id,omitempty"`
	Email    string                   `json:"email,omitempty" binding:"email"`
	Name     string                   `json:"name,omitempty"`
	Home     *GetArticlesReq_Location `json:"home,omitempty" binding:"required"`
	TestStr  string                   `json:"test_str,omitempty" binding:"required"`
}

type GetArticlesReply struct {
	Total    int64      `json:"total,omitempty" binding:"required"`
	Articles []*Article `json:"articles,omitempty"`
}

type Article struct {
	// 标题名称上注释
	Title    string `json:"title,omitempty"` // 标题名称右注释
	Content  string `json:"content,omitempty" binding:"required"`
	AuthorId int32  `json:"author_id,omitempty"`
}

type GetArticlesReq_Location struct {
	Lat float64 `json:"lat,omitempty" binding:"lte:90,gte:-90"`
	Lng float64 `json:"lng,omitempty" binding:"gte:-180,lte:180"`
}
