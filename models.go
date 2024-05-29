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

type Feed struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Name     string    `json:"name"`
	Url      string    `json:"url"`
	UserID   uuid.UUID `json:"user_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
		Name:     dbUser.Name,
	}
}


func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:       dbFeed.ID,
		CreateAt: dbFeed.CreateAt,
		UpdateAt: dbFeed.UpdateAt,
		Name:     dbFeed.Name,
		Url:      dbFeed.Url,
		UserID:   dbFeed.UserID,
	}
}