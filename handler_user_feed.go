package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AbdKaan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("follow command requires 1 argument: follow <url>")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("problem occured trying to get current user: %w", err)
	}
	userID := user.ID

	url := cmd.arguments[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("problem occured trying to get feed with url: %s. %w", url, err)
	}
	feedID := feed.ID

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	})
	if err != nil {
		return fmt.Errorf("problem occured trying to follow feed: %w", err)
	}

	fmt.Printf("Name: %s\nUser: %s\n", newFeedFollow.FeedName, newFeedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, _ command) error {
	feedFollows, err := s.db.GetFeedFollowsUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("error occured trying to get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("You aren't following any feeds yet. You can try using 'addfeed' command.")
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("- %s\n", feedFollow.FeedName)
	}

	return nil
}
