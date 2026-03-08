package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/AlfredBroda/gohair/storage"
	"github.com/AlfredBroda/gohair/storage/models"
	"github.com/gin-gonic/gin"
)

func NewArticleRouter(store *storage.Storage, ctx context.Context) *ArticleRouter {
	return &ArticleRouter{
		repository: *storage.NewArticleRepository(store, ctx),
	}
}

type ArticleRouter struct {
	repository storage.ArticleRepository
}

func (r *ArticleRouter) Register(engine *gin.Engine) {
	engine.GET("/a/:path", r.ArticleGet)
	engine.POST("/a/create", r.ArticleCreate)
	engine.DELETE("/a/:slug", r.ArticleDelete)
}

// ArticleGet handles the GET request to retrieve an article by its slug in JSON format
func (r *ArticleRouter) ArticleGet(c *gin.Context) {
	path := c.Param("path")
	log.Printf("Received request for article with path: %s", path)

	parts := strings.SplitN(path, ".", 2)
	slug := path
	ext := ""

	if len(parts) > 1 {
		slug = parts[0]
		ext = parts[len(parts)-1]
	}

	article, err := r.repository.GetArticleBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Article not found",
			"path":  c.Param("path"),
			"slug":  slug,
		})
		return
	}

	if ext == "html" {
		c.HTML(http.StatusOK, "articles/index.tmpl", gin.H{
			"title": article.Title,
			"body":  article.Content,
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
			"error":   "Invalid request body",
			"request": c.Request.Body,
		})
		return
	}

	err := r.repository.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create article: " + err.Error(),
			"slug":  article.Slug,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Article created successfully",
		"article": article,
	})
}

// ArticleDelete handles the DELETE request to delete an article by its slug
func (r *ArticleRouter) ArticleDelete(c *gin.Context) {
	slug := c.Param("slug")
	affected, err := r.repository.DeleteArticle(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Article could not be deleted: %s", err.Error()),
			"slug":  slug,
		})
		return
	}

	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Article not found",
			"slug":  slug,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":  "Article deleted successfully",
		"slug":     slug,
		"affected": affected,
	})
}
