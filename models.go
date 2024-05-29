package main

import (
	"time"

	"github.com/G0SU19O2/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Name     string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
		Name:     dbUser.Name,
	}
}