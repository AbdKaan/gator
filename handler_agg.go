package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, _ command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("problem fetching rss feed: %w", err)
	}

	fmt.Println(rssFeed)
	return nil
}
