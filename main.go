package main

import (
	"fmt"
	"log"
	"github.com/AgoCodeBro/gator/internal/config"
	"github.com/AgoCodeBro/gator/internal/database"
	"os"
	_ "github.com/lib/pq"
	"database/sql"
)


func main() {
	fmt.Printf("I love you kylah!\n")
	c, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to read config: %v\n", err)
	}
	
	db, err := sql.Open("postgres", c.DbURL)
	
	dbQueries := database.New(db)
	
	curState := &state{cfg : &c, db : dbQueries}
	
	cmds := commands{registeredCommands : make(map[string]func(*state, command) error)}
	registerCommands(&cmds)

	if len(os.Args) < 2 {
		log.Fatal("Please include a command")
	}

	//build command struct
	cmd := command{Name: os.Args[1], Args : os.Args[2:]}

	err = cmds.run(curState, cmd)
	if err != nil {
		log.Fatal(err)

	}

}

func registerCommands(cmds *commands) {	
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
}
