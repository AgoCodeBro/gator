package main

import ("github.com/AgoCodeBro/gator/internal/config"
				"github.com/AgoCodeBro/gator/internal/database"
				"github.com/google/uuid"
				"time"
				"fmt"
				"context"
				"errors"
				"log"
				_ "github.com/lib/pq"
				"os"
				"database/sql"
				)

func main() {
	// Creates the inital state struct
	cnfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
	}
	state := State{config : &cnfg}

	// Start up DB
	db, err := sql.Open("postgres", cnfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	fmt.Printf("%v", dbQueries)


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
		return errors.New("Login expects a username argument")
	}

	if err := st.config.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("User has been set\n")
	return nil
}

func registerHandler(st *State, cmd Command) (error) {
	if len(cmd.args) == 0 {
		return errors.New("Register expects a name argument")
	}
	
	// Check if user exists
	_, err := st.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
    // User already exists
    fmt.Println("User already exists")
    os.Exit(1)
	} else if err != sql.ErrNoRows {
    // Some other database error
    fmt.Printf("Database error: %v\n", err)
    os.Exit(1)
	}
	

	userParams := database.CreateUserParams{ID        : uuid.New(),
																					CreatedAt : time.Now(),
																					UpdatedAt : time.Now(),
																					Name      : cmd.args[0],
																					}

	user, err := st.db.CreateUser(context.Background(), userParams)
	fmt.Printf("Created user %v", user)
	if err != nil {
		return err
	}

	return nil

}
