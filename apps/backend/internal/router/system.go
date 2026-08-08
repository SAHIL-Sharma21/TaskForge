package router

import (
	"github.com/SAHIL-Sharma21/go-taskForge/internal/handler"
	"github.com/labstack/echo/v4"
)

func registerSystemRoutes(r *echo.Echo, h *handler.Handlers) {
	api := r.Group("/api")
	api.GET("/status", h.Health.CheckHealth)

	r.Static("/static", "static")

	r.GET("/docs", h.OpenAPI.ServeOpenAPIUI)
}
