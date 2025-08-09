package main

import (
	"github.com/AgoCodeBro/gator/internal/config"
	"github.com/AgoCodeBro/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
