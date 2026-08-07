package v1

import (
	"github.com/SAHIL-Sharma21/go-taskForge/internal/handler"
	"github.com/SAHIL-Sharma21/go-taskForge/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerCommentRouter(r *echo.Group, h *handler.CommentHandler, auth *middleware.AuthMiddleware) {
	// Comment operations
	comments := r.Group("/comments")
	comments.Use(auth.RequireAuth)

	// Individual comment operation
	dynamicComment := comments.Group("/:id")
	dynamicComment.PATCH("", h.UpdateComment)
	dynamicComment.DELETE("", h.DeleteComment)
}
