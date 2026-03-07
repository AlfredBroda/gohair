package transport

import (
	"net/http"

	"github.com/AlfredBroda/gohair/storage"
	"github.com/AlfredBroda/gohair/storage/models"
	"github.com/gin-gonic/gin"
)

func NewArticleRouter(dialector storage.Dialector) *ArticleRouter {
	return &ArticleRouter{dialector: dialector}
}

type ArticleRouter struct {
	dialector storage.Dialector
}

func (r *ArticleRouter) Register(engine *gin.Engine) {
	engine.GET("/a/:slug", r.ArticleGet)
	engine.POST("/a/create", r.ArticleCreate)
}

// ArticleGet handles the GET request to retrieve an article by its slug
func (r *ArticleRouter) ArticleGet(c *gin.Context) {
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

// ArticleCreate handles the POST request to create a new article
func (r *ArticleRouter) ArticleCreate(c *gin.Context) {
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := storage.CreateArticle(r.dialector, &article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create article",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Article created successfully",
		"article": article,
	})
}
