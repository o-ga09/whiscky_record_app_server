package entity

import (
	"time"
)

type UserID string

type User struct {
	User_ID UserID `json:"user_id" db:"user_id"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Modified_at time.Time `json:"modified_at" db:"modified_at"`
}