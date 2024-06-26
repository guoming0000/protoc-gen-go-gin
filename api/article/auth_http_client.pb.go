// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.3
// - protoc            v4.24.2
// source: api/auth.proto

package article

import (
	context "context"
	calloption "github.com/guoming0000/protoc-gen-go-gin/calloption"
	ecode "github.com/sunmi-OS/gocore/v2/api/ecode"
	http_request "github.com/sunmi-OS/gocore/v2/utils/http-request"
)

// AuthServiceHTTPClient is the client API for AuthService service.
type AuthServiceHTTPClient interface {
	// Push发送1 ||商米助手
	// 发送1<br>
	// 发送2<br>
	// > 发送3
	// 发送4
	// `import from json`
	// 这是一个链接 [Markdown语法](https://markdown.com.cn "最好的markdown教程")。
	//
	// | 错误码 | 错误消息 | 错误描述 |
	// | --- | ---- | ---- |
	// | 10207  | ath not match                     | 服务端token不一致      |
	// | 10224  | device not found                  | 设备不存在             |
	// | 10114  | task not found                    | 自动下载密钥任务不存在 |
	// | 10208  | binding key not found             | 绑定密钥不存在         |
	// | 10212  | no available license              | 没有可用license        |
	// | 10213  | license download times not enough | license 下载次数不足   |
	// | 500    |                                   | 程序异常               |
	Push(context.Context, *PushReq, ...calloption.CallOption) (*TResponse[PushReply], error)
	Pull(context.Context, *PushReq, ...calloption.CallOption) (*TResponse[RealResp], error)
	// 测试特殊的返回
	GetOneArticlePure(context.Context, *GetOneArticlePureReq, ...calloption.CallOption) (*GetOneArticlePureResp, error)
	PostOneArticlePure(context.Context, *GetOneArticlePureReq, ...calloption.CallOption) (*GetOneArticlePureResp, error)
}

type AuthServiceHTTPClientImpl struct {
	hh *http_request.HttpClient
}

func NewAuthServiceHTTPClient(hh *http_request.HttpClient) AuthServiceHTTPClient {
	return &AuthServiceHTTPClientImpl{hh: hh}
}

func (c *AuthServiceHTTPClientImpl) Push(ctx context.Context, req *PushReq, opts ...calloption.CallOption) (*TResponse[PushReply], error) {
	resp := &TResponse[PushReply]{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/private/v1/push")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(int(resp.Code), resp.Msg)
	}
	return resp, err
}

func (c *AuthServiceHTTPClientImpl) Pull(ctx context.Context, req *PushReq, opts ...calloption.CallOption) (*TResponse[RealResp], error) {
	resp := &TResponse[RealResp]{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/private/pull")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(int(resp.Code), resp.Msg)
	}
	return resp, err
}

func (c *AuthServiceHTTPClientImpl) GetOneArticlePure(ctx context.Context, req *GetOneArticlePureReq, opts ...calloption.CallOption) (*GetOneArticlePureResp, error) {
	// TODO: GET method not support
	return nil, ecode.NewV2(-1, "GET method not support")
}
func (c *AuthServiceHTTPClientImpl) PostOneArticlePure(ctx context.Context, req *GetOneArticlePureReq, opts ...calloption.CallOption) (*GetOneArticlePureResp, error) {
	resp := &GetOneArticlePureResp{}
	r := c.hh.Client.R().SetContext(ctx)
	for _, opt := range opts {
		opt(r)
	}
	_, err := r.SetBody(req).SetResult(resp).Post("/v1/get/article")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
