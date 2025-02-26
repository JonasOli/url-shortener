package model

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Salt       string    `json:"salt"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
}
