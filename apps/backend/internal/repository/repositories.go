package repository

import "github.com/SAHIL-Sharma21/go-taskForge/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
