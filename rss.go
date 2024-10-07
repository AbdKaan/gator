package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("problem occured creating request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("problem occured doing the request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("problem occured reading response body: %w", err)
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return &RSSFeed{}, fmt.Errorf("problem occured unmarshaling data: %w", err)
	}

	// Decode escaped HTML entities (like &ldquo;) on title and description fields
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return &rssFeed, nil
}
