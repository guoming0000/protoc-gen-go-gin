package article

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sunmi-OS/gocore/v2/api"
	"github.com/sunmi-OS/gocore/v2/utils"
	http_request "github.com/sunmi-OS/gocore/v2/utils/http-request"
)

var ErrIncorrectCode = errors.New("incorrect http status")

func MustCode200(cli *resty.Client, resp *resty.Response) error {
	if resp.StatusCode() == http.StatusOK {
		return fmt.Errorf("%w:%d", ErrIncorrectCode, resp.StatusCode())
	}
	return nil
}

var (
	PrivateHosts = map[string]string{
		"onl":   "http://vpc-eco-iotlink-gateway.sunmi.com",
		"uat":   "http://vpc-eco-iotlink-gateway.uat.sunmi.com",
		"test":  "http://vpc-eco-iotlink-gateway.test.sunmi.com",
		"dev":   "http://vpc-eco-iotlink-gateway.dev.sunmi.com",
		"local": "http://localhost:20080",
	}
	PublicHosts = map[string]string{
		"onl":   "http://api.sunmi.com",
		"uat":   "http://api.uat.sunmi.com",
		"test":  "http://api.test.sunmi.com",
		"dev":   "http://api.dev.sunmi.com",
		"local": "http://localhost:20080",
	}
)

func GetHostUrl(isPrivate bool, envStrs ...string) string {
	envStr := utils.GetRunTime()
	if len(envStrs) > 0 {
		envStr = envStrs[0]
	}
	hostMap := PublicHosts
	if isPrivate {
		hostMap = PrivateHosts
	}
	var ok bool
	var hostUrl string
	if hostUrl, ok = hostMap[envStr]; ok {
		return hostUrl
	}
	hostUrl = hostMap["onl"]
	return hostUrl
}

// 根据需要配置client相关参数
func newClient(isPrivate bool, strs ...string) *http_request.HttpClient {
	str := "onl"
	if len(strs) > 0 {
		str = strs[0]
	}
	h := http_request.New()
	h.Client.SetTimeout(time.Second * 30).OnAfterResponse(MustCode200)
	h.Client.SetHostURL(GetHostUrl(isPrivate, str))
	appID := os.Getenv("APP_ID")
	if appID == "" {
		appID = "default-iotlink-go"
	}
	h.Client.SetHeader("User-Agent", appID) // 从环境变量取
	return h
}

// NewClient 使用者也可以根据需要，自己生成newClient方法，然后调用NewBlogHTTPClient方法
func NewClient(strs ...string) BlogServiceHTTPClient {
	return NewBlogServiceHTTPClient(newClient(true, strs...))
}

func NewPublicClient(strs ...string) BlogServiceHTTPClient {
	return NewBlogServiceHTTPClient(newClient(false, strs...))
}

var DefaultClient = NewClient()
var PubClient = NewPublicClient()

func test() {
	ctx := &api.Context{}
	resp, err := DefaultClient.GetArticles(ctx, &GetArticlesReq{})
	_ = resp
	_ = err
}
