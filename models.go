package main

import (
	"database/sql"
	"time"

	"github.com/Saakhr/rssaggr/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseToUser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		ApiKey:    dbuser.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID    `json:"id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Name          string       `json:"name"`
	Url           string       `json:"url"`
	UserID        uuid.UUID    `json:"user_id"`
	LastFetchedAt sql.NullTime `json:"last_fetched_at"`
}

func databaseToFeed(dfeed database.Feed) Feed {
	return Feed{
		ID:        dfeed.ID,
		CreatedAt: dfeed.CreatedAt,
		UpdatedAt: dfeed.UpdatedAt,
		Name:      dfeed.Name,
		Url:       dfeed.Url,
		UserID:    dfeed.UserID,
	}
}
func databaseToFeeds(dfeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dfeed {
		feeds = append(feeds, databaseToFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseToFeedFollow(dfeedfollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dfeedfollow.ID,
		CreatedAt: dfeedfollow.CreatedAt,
		UpdatedAt: dfeedfollow.UpdatedAt,
		UserID:    dfeedfollow.UserID,
		FeedID:    dfeedfollow.FeedID,
	}
}
func databaseToFeedFollows(dfeedfollows []database.FeedFollow) []FeedFollow {
	follows := []FeedFollow{}
	for _, follow := range dfeedfollows {
		follows = append(follows, databaseToFeedFollow(follow))
	}
	return follows
}

type Post struct {
	ID          uuid.UUID      `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	PublishedAt time.Time      `json:"published_at"`
	Url         string         `json:"url"`
	FeedID      uuid.UUID      `json:"feed_id"`
}

func databaseToPost(dpost database.Post) Post {
	return Post{
		ID:          dpost.ID,
		CreatedAt:   dpost.CreatedAt,
		UpdatedAt:   dpost.UpdatedAt,
		Title:       dpost.Title,
		Description: dpost.Description,
		PublishedAt: dpost.PublishedAt,
		Url:         dpost.Url,
		FeedID:      dpost.FeedID,
	}
}

func databaseToPosts(dposts []database.Post) []Post {
	posts := []Post{}
	for _, post := range dposts {
		posts = append(posts, databaseToPost(post))
	}
	return posts
}
