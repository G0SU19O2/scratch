// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID       uuid.UUID
	CreateAt time.Time
	UpdateAt time.Time
	Name     string
	Url      string
	UserID   uuid.UUID
}

type FeedFollow struct {
	ID       uuid.UUID
	CreateAt time.Time
	UpdateAt time.Time
	FeedID   uuid.UUID
	UserID   uuid.UUID
}

type User struct {
	ID       uuid.UUID
	CreateAt time.Time
	UpdateAt time.Time
	Name     string
	ApiKey   string
}
