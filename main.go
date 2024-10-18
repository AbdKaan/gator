package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AbdKaan/gator/internal/config"
	"github.com/AbdKaan/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error trying to read config: %v", err)
	}

	dbURL := "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	config_state := state{dbQueries, &cfg}

	var handlerFunctions = make(map[string]func(*state, command) error)
	cli_commands := commands{handlerFunctions}

	cli_commands.register("login", handlerLogin)
	cli_commands.register("register", handlerRegister)
	cli_commands.register("reset", handlerReset)
	cli_commands.register("users", handlerListUsers)
	cli_commands.register("agg", handlerAgg)
	cli_commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cli_commands.register("feeds", handlerListFeeds)
	cli_commands.register("follow", middlewareLoggedIn(handlerFollowFeed))
	cli_commands.register("following", handlerFollowing)
	cli_commands.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	cli_commands.register("browse", middlewareLoggedIn(handlerBrowse))

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("not enough arguments given")
		os.Exit(1)
	}

	cmd := command{arguments[1], arguments[2:]}
	err = cli_commands.run(&config_state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
