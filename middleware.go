package main

import (
	"context"
	"fmt"

	"github.com/AbdKaan/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			return fmt.Errorf("couldn't get the current user %s with error: %w", s.cfg.Current_user_name, err)
		}

		return handler(s, cmd, user)
	}
}
