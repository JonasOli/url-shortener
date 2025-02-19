package model

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
}
