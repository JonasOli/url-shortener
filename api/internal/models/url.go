package models

import "time"

type URL struct {
	ID          int       `json:"id"`
	Original    string    `json:"original"`
	Short       string    `json:"short"`
	Visit_count int       `json:"visit_count"`
	Created_at  time.Time `json:"created_at"`
}
