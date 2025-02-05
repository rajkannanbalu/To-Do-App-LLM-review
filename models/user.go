package models

import "time"

// User represent the user model
type UserDB struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	LastActive time.Time `json:"last_active"`
}
