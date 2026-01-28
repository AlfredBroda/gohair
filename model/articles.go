package model

import (
	"context"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

func GetArticleBySlug(dialector gorm.Dialector, slug string) (*Article, error) {
	db, err := InitDB(dialector)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	// Read
	article, err := gorm.G[Article](db).Where("slug = ?", slug).First(ctx) // find Article with integer primary key

	return &article, err
}
