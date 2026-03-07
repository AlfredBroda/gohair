package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}
