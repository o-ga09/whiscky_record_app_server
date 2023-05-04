package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserID int64

type User struct {
	ID UserID `json:"id" db:"id"`
	User_ID string `json:"user_id" db:"user_id"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Modified_at time.Time `json:"modified_at" db:"modified_at"`
}

func (u *User) ComparePassword(uid string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.User_ID),[]byte(uid))
}