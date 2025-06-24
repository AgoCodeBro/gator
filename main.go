package main

import ("github.com/AgoCodeBro/gator/internal/config"
				"fmt"
				"errors"
				"log"
				"os"
				)

func main() {
	// Creates the inital state struct
	cnfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
	}
	state := State{config : &cnfg}

	// Creates commands struct and adds handlers
	cmds := Commands{commandsMap : make(map[string]func(*State, Command)(error))}
	cmds.register("login", handlerLogin)


	// Reads given arguments
	if len(os.Args) < 2 {
		err := errors.New("No command given")
		log.Fatal(err)
	}

	cmdName := os.Args[1]
	var cmdArgs []string
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}

	cmd := Command{name : cmdName,
								args : cmdArgs,
								}
	err = cmds.run(&state, cmd)
	if err != nil {
		log.Fatal(err)
	}

}

func handlerLogin(st *State, cmd Command) (error) {
	if len(cmd.args) == 0 {
		return errors.New("Login expect a username argument")
	}

	if err := st.config.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("User has been set\n")
	return nil
}
