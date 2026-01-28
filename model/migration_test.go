package model

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
	dialector := ConfigureMySQL(DBConfig{
		DBUser: "root",
		DBPass: "password",
		DBAddr: "localhost",
		DBPort: 3306,
	})
	db, err := InitDB(dialector)
	require.NoError(t, err)

	err = Migrate(dialector)
	require.NoError(t, err)

	ctx := context.Background()
	// Create
	err = gorm.G[Article](db).Create(ctx, &Article{
		Slug:    "article-1",
		Title:   "First Article",
		Summary: "This is the summary",
		Content: "<p>This is the content</p>",
	})
	require.NoError(t, err)

	res, err := GetArticleBySlug(dialector, "article-1")

	require.NoError(t, err)
	require.NotEmpty(t, res)
}
