package main

import (
	"fmt"
	"log"
	"github.com/AgoCodeBro/gator/internal/config"
	"os"
)


func main() {
	fmt.Printf("I love you kylah!\n")
	c, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
	}
	curState := &state{cfg : &c}
	
	cmds := commands{registeredCommands : make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Please include a command")
	}

	//build command struct
	cmd := command{Name: os.Args[1], Args : os.Args[2:]}

	err = cmds.run(curState, cmd)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("Database: %v\nUser: %v\n", c.DbURL, c.CurrentUserName)
}
