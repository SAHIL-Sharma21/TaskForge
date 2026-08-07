package v1

import (
	"github.com/SAHIL-Sharma21/go-taskForge/internal/handler"
	"github.com/SAHIL-Sharma21/go-taskForge/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerTodoRouter(r *echo.Group, h *handler.TodoHandler, ch *handler.CommentHandler, auth *middleware.AuthMiddleware) {
	// TODO Operation

	todos := r.Group("/todos")
	todos.Use(auth.RequireAuth)

	// Collection operations
	todos.POST("", h.CreateTodo)
	todos.GET("", h.GetTodos)
	todos.GET("/stats", h.GetTodoStats)

	// Individual todo operations
	dynamicTodo := todos.Group("/:id")
	dynamicTodo.GET("", h.GetTodoByID)
	dynamicTodo.PATCH("", h.UpdateTodo)
	dynamicTodo.DELETE("", h.DeleteTodo)

	// TODO Comments
	todoComments := dynamicTodo.Group("/comments")
	todoComments.POST("", ch.AddComment)
	todoComments.GET("", ch.GetCommentsByTodoID)
}
