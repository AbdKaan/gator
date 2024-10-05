package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username is required")
	}

	err := s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set.")

	return nil
}
