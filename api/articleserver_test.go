package api

import (
	"github.com/gin-gonic/gin"
	"github.com/guoming0000/protoc-gen-go-gin/api/v1/article"
	"github.com/sunmi-OS/gocore/v2/api"
)

func InitServerTest() {
	router := gin.Default()

	httpSrv := NewBlogServiceHttpServer()
	article.RegisterBlogServiceHTTPServer(router, httpSrv)
}

func NewBlogServiceHttpServer() article.BlogServiceHTTPServer {
	return &blogServiceHttpServer{}
}

type blogServiceHttpServer struct {
	// 各种db的定义
}

func (b *blogServiceHttpServer) GetArticles(context *api.Context, req *article.GetArticlesReq) (*article.GetArticlesReply, error) {
	req.Validate()
	return &article.GetArticlesReply{
		Total: 775755,
	}, nil
}

func (b *blogServiceHttpServer) CreateArticle(context *api.Context, a *article.Article) (*article.Article, error) {
	return &article.Article{
		Title:    "6666",
		Content:  "content",
		AuthorId: 1,
	}, nil
}
