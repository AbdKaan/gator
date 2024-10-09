package main

import "github.com/AbdKaan/gator/internal/database"

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

}
