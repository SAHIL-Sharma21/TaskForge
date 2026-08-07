package v1

import (
	"github.com/SAHIL-Sharma21/go-taskForge/internal/handler"
	"github.com/SAHIL-Sharma21/go-taskForge/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers, middleware *middleware.Middlewares) {
	// Register todo routes
	registerTodoRouter(router, handlers.Todo, handlers.Comment, middleware.Auth)

	// Register category routes
	registerCategoryRouter(router, handlers.Category, middleware.Auth)

	// Register comment routes
	registerCommentRouter(router, handlers.Comment, middleware.Auth)
}
