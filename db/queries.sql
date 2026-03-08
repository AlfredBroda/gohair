-- name: GetArticleBySlug :one
SELECT id, slug, title, summary, content, created_at, updated_at, deleted_at
FROM articles
WHERE slug = ?
  AND deleted_at IS NULL;

-- name: InsertArticle :exec
INSERT INTO articles (slug, title, summary, content, created_at, updated_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- name: GetArticleByID :one
SELECT id, slug, title, summary, content, created_at, updated_at, deleted_at
FROM articles
WHERE id = LAST_INSERT_ID();

-- name: DeleteArticle :exec
UPDATE articles
SET deleted_at = CURRENT_TIMESTAMP, slug = null
WHERE slug = ?;

-- name: UpdateArticle :exec
UPDATE articles
SET title = ?, summary = ?, content = ?
WHERE slug = ? AND deleted_at IS NULL;

-- name: GetArticleBySlugForUpdate :one
SELECT id, slug, title, summary, content, created_at, updated_at, deleted_at
FROM articles
WHERE slug = ? AND deleted_at IS NULL;
