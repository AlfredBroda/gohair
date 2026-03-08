package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlfredBroda/gohair/storage/models"
)

func NewArticleRepository(storage *Storage, ctx context.Context) *ArticleRepository {
	return &ArticleRepository{
		queries: storage.Queries,
		context: ctx,
	}
}

type ArticleRepository struct {
	queries *Queries
	context context.Context
}

func (r *ArticleRepository) GetArticleBySlug(slug string) (*models.Article, error) {
	// sqlc-generated method accepts sql.NullString
	rec, err := r.queries.GetArticleBySlug(r.context, sql.NullString{String: slug, Valid: true})
	if err != nil {
		return nil, err
	}
	// convert storage.Article (generated) to API model
	var deletedAt *time.Time
	if rec.DeletedAt.Valid {
		deletedAt = &rec.DeletedAt.Time
	}
	return &models.Article{
		ID:        uint(rec.ID),
		Slug:      rec.Slug.String,
		Title:     rec.Title,
		Summary:   rec.Summary.String,
		Content:   rec.Content.String,
		CreatedAt: rec.CreatedAt.Time,
		UpdatedAt: rec.UpdatedAt.Time,
		DeletedAt: deletedAt,
	}, nil
}

func (r *ArticleRepository) CreateArticle(article *models.Article) error {
	// execute insert then fetch the newly created row by ID
	if err := r.queries.InsertArticle(r.context, InsertArticleParams{
		Slug:    sql.NullString{String: article.Slug, Valid: true},
		Title:   article.Title,
		Summary: sql.NullString{String: article.Summary, Valid: true},
		Content: sql.NullString{String: article.Content, Valid: true},
	}); err != nil {
		return err
	}
	rec, err := r.queries.GetArticleByID(r.context)
	if err != nil {
		return err
	}
	// convert returned record back to models.Article
	var deletedAt *time.Time
	if rec.DeletedAt.Valid {
		deletedAt = &rec.DeletedAt.Time
	}
	article.ID = uint(rec.ID)
	article.Slug = rec.Slug.String
	article.Title = rec.Title
	article.Summary = rec.Summary.String
	article.Content = rec.Content.String
	article.CreatedAt = rec.CreatedAt.Time
	article.UpdatedAt = rec.UpdatedAt.Time
	article.DeletedAt = deletedAt
	return nil
}

func (r *ArticleRepository) DeleteArticle(slug string) (int64, error) {
	res, err := r.queries.db.ExecContext(r.context,
		"UPDATE articles SET deleted_at = CURRENT_TIMESTAMP, slug = NULL WHERE slug = ?", slug)
	if err != nil {
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
