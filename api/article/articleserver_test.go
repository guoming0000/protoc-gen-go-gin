package article

import (
	"github.com/gin-gonic/gin"
	"github.com/sunmi-OS/gocore/v2/api"
)

func InitServerTest() {
	router := gin.Default()

	httpSrv := NewBlogServiceHttpServer()
	RegisterBlogServiceHTTPServer(router, httpSrv)
}

func NewBlogServiceHttpServer() BlogServiceHTTPServer {
	return &blogServiceHttpServer{}
}

type blogServiceHttpServer struct {
	// 各种db的定义
}

func (b *blogServiceHttpServer) GetArticles(context *api.Context, req *GetArticlesReq) (*GetArticlesReply, error) {
	return &GetArticlesReply{
		Total: 775755,
	}, nil
}

func (b *blogServiceHttpServer) CreateArticle(context *api.Context, a *Article) (*Article, error) {
	return &Article{
		Title:    "6666",
		Content:  "content",
		AuthorId: 1,
	}, nil
}
