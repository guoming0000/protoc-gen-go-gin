// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v0.0.2
// - protoc            v4.24.2
// source: api/auth.proto

package article

import (
	gin "github.com/gin-gonic/gin"
	api "github.com/sunmi-OS/gocore/v2/api"
)

// AuthServiceHTTPServer is the server API for AuthService service.
type AuthServiceHTTPServer interface {
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
	Push(*api.Context, *PushReq) (*PushReply, error)
	Pull(*api.Context, *PushReq) (*RealResp, error)
}

func RegisterAuthServiceHTTPServer(s *gin.Engine, srv AuthServiceHTTPServer) {
	r := s.Group("/")
	r.POST("/private/v1/push", _AuthService_Push_HTTP_Handler(srv))  // Push发送1 ||商米助手
	r.POST("/private/v1/push2", _AuthService_Push_HTTP_Handler(srv)) // Push发送1 ||商米助手
	r.POST("/private/push", _AuthService_Push_HTTP_Handler(srv))     // Push发送1 ||商米助手
	r.POST("/private/pull", _AuthService_Pull_HTTP_Handler(srv))
}

func _AuthService_Push_HTTP_Handler(srv AuthServiceHTTPServer) func(g *gin.Context) {
	return func(g *gin.Context) {
		req := &PushReq{}
		var err error
		ctx := api.NewContext(g)
		err = parseReq(g, &ctx, req)
		err = checkValidate(err)
		if err != nil {
			setRetJSON(&ctx, nil, err)
			return
		}
		resp, err := srv.Push(&ctx, req)
		setRetJSON(&ctx, resp, err)
	}
}

func _AuthService_Pull_HTTP_Handler(srv AuthServiceHTTPServer) func(g *gin.Context) {
	return func(g *gin.Context) {
		req := &PushReq{}
		var err error
		ctx := api.NewContext(g)
		err = parseReq(g, &ctx, req)
		err = checkValidate(err)
		if err != nil {
			setRetJSON(&ctx, nil, err)
			return
		}
		resp, err := srv.Pull(&ctx, req)
		setRetJSON(&ctx, resp, err)
	}
}
