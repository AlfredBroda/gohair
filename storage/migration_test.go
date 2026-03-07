package storage_test

import (
	"context"
	"testing"

	"github.com/AlfredBroda/gohair/storage"
	"github.com/AlfredBroda/gohair/storage/models"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
	dialector := storage.ConfigureMySQL(storage.DBConfig{
		DBUser: "root",
		DBPass: "password",
		DBAddr: "localhost",
		DBPort: 3306,
	})
	db, err := storage.InitDB(dialector)
	require.NoError(t, err)

	err = storage.Migrate(dialector)
	require.NoError(t, err)

	ctx := context.Background()
	// Create
	err = gorm.G[models.Article](db).Create(ctx, &models.Article{
		Slug:    "article-1",
		Title:   "First Article",
		Summary: "This is the summary",
		Content: "<p>This is the content</p>",
	})
	require.NoError(t, err)

	res, err := storage.GetArticleBySlug(dialector, "article-1")

	require.NoError(t, err)
	require.NotEmpty(t, res)
}
