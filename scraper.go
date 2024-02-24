package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Saakhr/rssaggr/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurency int, timeBetweenReq time.Duration) {
	log.Printf("Scrapping on %v goroutines every %s Duration", concurency, timeBetweenReq)
	ticker := time.NewTicker(timeBetweenReq)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurency))
		if err != nil {
			log.Println("error fetching feeds")
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}
func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched: ", err)
		return
	}
	rssfeed, err := UrlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed data from url: ", err)
		return
	}
	for _, item := range rssfeed.Channel.Item {
		var published time.Time
		if item.PubDate == "" {
			published, err = parseDateWithFormats(item.Date)
		} else {
			published, err = parseDateWithFormats(item.PubDate)
		}
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: published,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to push feed to db with err: ", err)
		}
	}
	log.Printf("Feed %s collected %v posts found", feed.Name, len(rssfeed.Channel.Item))

}
