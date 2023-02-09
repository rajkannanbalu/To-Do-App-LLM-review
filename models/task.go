package models

import (
	
	"time"
)

// TaskDB represent the task model
type TaskDB struct {
	ID        int64      `json:"id"`
	Name   	  string     `json:"name"`
	Status    string     `json:"status"`
	Comment   string     `json:"comment"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	UserID    int64      `json:"user_id"`
}