package user

import (
	"time"
)

type User struct {
	Id           int
	Name         string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
