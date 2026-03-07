package storage

import (
	"context"

	"github.com/AlfredBroda/gohair/storage/models"
	"gorm.io/gorm"
)

func GetArticleBySlug(dialector gorm.Dialector, slug string) (*models.Article, error) {
	db, err := InitDB(dialector)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	// find Article with string primary key
	article, err := gorm.G[models.Article](db).Where("slug = ?", slug).First(ctx)

	return &article, err
}
