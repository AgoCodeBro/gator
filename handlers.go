package main

import(
	"errors"
	"fmt"
	"github.com/AgoCodeBro/gator/internal/config"
	"context"
	"github.com/AgoCodeBro/gator/internal/database"
	"github.com/google/uuid"
	"time"
	"database/sql"
)

func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) == 0 {
		return errors.New("Login expects a username as an argument")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Failed to find user %v: %v", cmd.Args[0], err)
	}

	s.cfg.CurrentUserName = cmd.Args[0]
	err = config.SetUser(s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to set the User: %v", err)
	}

	fmt.Printf("%v has been set as the current user\n", s.cfg.CurrentUserName)

	return nil
}



func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Register expects a username as an argument")
	}

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, cmd.Args[0])
	if err == nil {
		return fmt.Errorf("User %v already exists", cmd.Args[0])
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("Failed to register %v: %v", cmd.Args[0], err)
	}

	args := database.CreateUserParams { 
		ID : uuid.New(),
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		Name :  cmd.Args[0],
	}
	
	_, err = s.db.CreateUser(ctx, args)
	if err != nil {
		return fmt.Errorf("Error while adding %v to the database: %v", cmd.Args[0], err)
	}
	
	s.cfg.CurrentUserName = cmd.Args[0]
	err = config.SetUser(s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to set the User: %v", err)
	}

	fmt.Printf("%v has been set as the current user\n", s.cfg.CurrentUserName)

	return nil
	
}
	
func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Error occured while trying to reset the users table: %v", err)
	}

	return nil
}
