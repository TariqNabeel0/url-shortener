package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	OriginalURL string `gorm:"not null"`
	ShortCode string `gorm:"uniqueIndex"`
}
