// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.3
// - protoc            v4.24.2
// source: api/article.proto

package article

import (
	sonic "github.com/bytedance/sonic"
)

func (m *GetArticlesReq) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *GetArticlesReq) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *GetArticlesReq) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *GetArticlesReq) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}

func (m *GetArticlesReply) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *GetArticlesReply) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *GetArticlesReply) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *GetArticlesReply) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}

func (m *Article) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *Article) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *Article) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *Article) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}
