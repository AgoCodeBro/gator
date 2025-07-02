package main

import ("github.com/AgoCodeBro/gator/internal/config"
				"github.com/AgoCodeBro/gator/internal/database"
				)

type State struct {
	config *config.Config
	db     *database.Queries
}
