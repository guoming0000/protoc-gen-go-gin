// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v0.0.2
// - protoc            v3.21.12
// source: api/article/article.proto

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
	m = new(GetArticlesReq)
	err := sonic.Unmarshal(buf, m)
	return err
}

func (m *GetArticlesReq) UnmarshalString(str string) error {
	m = new(GetArticlesReq)
	err := sonic.UnmarshalString(str, m)
	return err
}

func (m *GetArticlesReply) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *GetArticlesReply) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *GetArticlesReply) Unmarshal(buf []byte) error {
	m = new(GetArticlesReply)
	err := sonic.Unmarshal(buf, m)
	return err
}

func (m *GetArticlesReply) UnmarshalString(str string) error {
	m = new(GetArticlesReply)
	err := sonic.UnmarshalString(str, m)
	return err
}

func (m *Article) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *Article) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *Article) Unmarshal(buf []byte) error {
	m = new(Article)
	err := sonic.Unmarshal(buf, m)
	return err
}

func (m *Article) UnmarshalString(str string) error {
	m = new(Article)
	err := sonic.UnmarshalString(str, m)
	return err
}
