package service

import (
	"github.com/SAHIL-Sharma21/go-taskForge/internal/lib/job"
	"github.com/SAHIL-Sharma21/go-taskForge/internal/repository"
	"github.com/SAHIL-Sharma21/go-taskForge/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
