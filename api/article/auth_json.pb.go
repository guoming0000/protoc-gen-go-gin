// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.2
// - protoc            v4.24.2
// source: api/auth.proto

package article

import (
	sonic "github.com/bytedance/sonic"
)

func (m *PushReq) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *PushReq) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *PushReq) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *PushReq) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}

func (m *PushReply) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *PushReply) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *PushReply) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *PushReply) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}

func (m *RealResp) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *RealResp) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *RealResp) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *RealResp) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}

func (m *UploadOssdk3RdReq) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *UploadOssdk3RdReq) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *UploadOssdk3RdReq) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *UploadOssdk3RdReq) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}

func (m *GetOneArticlePureReq) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *GetOneArticlePureReq) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *GetOneArticlePureReq) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *GetOneArticlePureReq) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}

func (m *GetOneArticlePureResp) Marshal() ([]byte, error) {
	return sonic.Marshal(m)
}

func (m *GetOneArticlePureResp) MarshalString() (string, error) {
	return sonic.MarshalString(m)
}

func (m *GetOneArticlePureResp) Unmarshal(buf []byte) error {
	return sonic.Unmarshal(buf, m)
}

func (m *GetOneArticlePureResp) UnmarshalString(str string) error {
	return sonic.UnmarshalString(str, m)
}
