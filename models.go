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

type FeedFollow struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	FeedID   uuid.UUID `json:"feed_id"`
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

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = databaseFeedToFeed(dbFeed)
	}
	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:       dbFeedFollow.ID,
		CreateAt: dbFeedFollow.CreateAt,
		UpdateAt: dbFeedFollow.UpdateAt,
		FeedID:   dbFeedFollow.FeedID,
		UserID:   dbFeedFollow.UserID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := make([]FeedFollow, len(dbFeedFollows))
	for i, dbFeedFollow := range dbFeedFollows {
		feedFollows[i] = databaseFeedFollowToFeedFollow(dbFeedFollow)
	}
	return feedFollows
}