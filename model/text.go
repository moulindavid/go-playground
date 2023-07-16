package model

import (
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}
