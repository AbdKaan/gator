package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/AbdKaan/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) > 1 {
		return errors.New("wrong usage. please try: browse <limit>\n limit is optional and is the number of posts, default is 2")
	}

	var limit int32
	if len(cmd.arguments) == 1 {
		limitArg, err := strconv.ParseInt(cmd.arguments[0], 10, 32)
		if err != nil {
			return fmt.Errorf("problem occured parsing the 'limit' argument: %w", err)
		}
		limit = int32(limitArg)
	} else {
		limit = 2
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("problem occured retrieving posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.Name)
		fmt.Printf("--- %s ---\n", post.Title)
		// Descriptions are scuffed
		//fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
