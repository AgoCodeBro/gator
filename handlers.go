package main

import(
	"errors"
	"fmt"
	"github.com/AgoCodeBro/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) == 0 {
		return errors.New("Login expects a username as an argument")
	}

	s.cfg.CurrentUserName = cmd.Args[0]
	err := config.SetUser(s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	fmt.Printf("%v has been set as the current user\n", s.cfg.CurrentUserName)

	return nil
}
