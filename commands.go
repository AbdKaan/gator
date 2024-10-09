package main

import (
	"errors"

	"github.com/AbdKaan/gator/internal/config"
	"github.com/AbdKaan/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.cmds[cmd.name]; !ok {
		return errors.New("given command doesn't exist")
	}

	err := c.cmds[cmd.name](s, cmd)
	return err
}
