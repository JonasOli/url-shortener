package models

import (
	"time"
)

type URL struct {
	OriginalURL string    `bson:"original_url"`
	ShortCode   string    `bson:"short_code"`
	CreatedAt   time.Time `bson:"created_at"`
	ClickCount  int64     `bson:"click_count"`
}
