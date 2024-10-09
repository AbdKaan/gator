package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AbdKaan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return errors.New("wrong usage. please try: addfeed <name> <url>")
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("couldn't get the current user %s with error: %w", s.cfg.Current_user_name, err)
	}
	user_id := user.ID

	new_feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user_id,
	})
	if err != nil {
		return fmt.Errorf("couldn't add the feed: %w", err)
	}

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user_id,
		FeedID:    new_feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't add the feedFollow: %v, with error: %w", newFeedFollow, err)
	}

	fmt.Println("Feed has been added.")
	printFeedInfo(new_feed)

	return nil
}

func printFeedInfo(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsAndUserName(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	for i, feed := range feeds {
		fmt.Printf("%d -\n", i+1)
		fmt.Printf("    Name:     %s\n", feed.Name)
		fmt.Printf("    URL:      %s\n", feed.Url)
		fmt.Printf("    Username: %s\n", feed.UserName)
	}

	return nil
}
