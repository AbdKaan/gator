package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/AbdKaan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("wrong usage. please try: agg <duration>\n example usage: agg 1h5m46s")
	}

	durationStr := cmd.arguments[0]

	timeBetweenRequests, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("problem occured parsing duration: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error occured trying to scrape feeds: %w", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("couldn't get the current user %s with error: %w", s.cfg.Current_user_name, err)
	}
	user_id := user.ID
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background(), user_id)
	if err != nil {
		return fmt.Errorf("problem fetching next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("problem marking fetched feed: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("problem fetching rss feed: %w", err)
	}

	fmt.Printf("Feed: %s, got %d Titles\n", nextFeed.Name, len(rssFeed.Channel.Item))
	for _, feed := range rssFeed.Channel.Item {
		timePubDate, err := time.Parse(time.RFC1123, feed.PubDate)
		if err != nil {
			return fmt.Errorf("problem occured parsing publish date: %w", err)
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       feed.Title,
			Url:         feed.Link,
			Description: sql.NullString{String: feed.Description, Valid: true},
			PublishedAt: timePubDate,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("problem occured creating post: %v", err)
			continue
		}
	}

	return nil
}
