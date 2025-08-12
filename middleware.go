package main

import (
	"fmt"
	"context"
	"github.com/AgoCodeBro/gator/internal/database"
)

func middlewareLoggedIn(handler func (s *state, cmd command, user database.User) error) func (*state, command) error {
	normalHandler := func (s *state, cmd command) error {
		userName := s.cfg.CurrentUserName
		user, err := s.db.GetUser(context.Background(), userName)
		if err != nil {
			return fmt.Errorf("Failed to find the current user")
		}

		return handler(s, cmd, user)
	}

	return normalHandler
}


