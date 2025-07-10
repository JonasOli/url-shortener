package models

import (
	"time"
)

type URL struct {
	ID          uint64     `json:"id" gorm:"primaryKey"`
	OriginalURL string     `json:"original_url" gorm:"not null;index" binding:"required,url"`
	ShortCode   string     `json:"short_code" gorm:"uniqueIndex;not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	ClickCount  int64      `json:"click_count" gorm:"default:0"`
}
