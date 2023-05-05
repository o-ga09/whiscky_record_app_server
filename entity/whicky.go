package entity

import (
	"time"
)

type WhickyRecord struct {
    UserID   string      `db:"user_id" json:"userId"`
    Name     string     `db:"whisky_name" json:"name"`
    DrankAt  time.Time `db:"drankAt" json:"drankAt"`
    ImageURL string     `db:"imageUrl" json:"imageUrl"`
}