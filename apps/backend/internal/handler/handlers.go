package handler

import (
	"github.com/SAHIL-Sharma21/go-taskForge/internal/server"
	"github.com/SAHIL-Sharma21/go-taskForge/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
