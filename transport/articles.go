package transport

import (
	"net/http"

	"github.com/AlfredBroda/gohair/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewArticleRouter(dialector storage.Dialector) *ArticleRouter {
	return &ArticleRouter{dialector: dialector}
}

type ArticleRouter struct {
	dialector gorm.Dialector
}

func (r *ArticleRouter) Register(engine *gin.Engine) {
	engine.GET("/a/:slug", r.ArticleHandler)
}

func (r *ArticleRouter) ArticleHandler(c *gin.Context) {
	slug := c.Param("slug")
	article, err := storage.GetArticleBySlug(r.dialector, slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Article not found",
			"slug":  slug,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}
