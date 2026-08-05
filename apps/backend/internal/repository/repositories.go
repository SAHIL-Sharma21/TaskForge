package repository

import "github.com/SAHIL-Sharma21/go-taskForge/internal/server"

type Repositories struct {
	Todo    *TodoRepository
	comment *CommentRepository
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{
		Todo:    NewTodoRepository(s),
		comment: NewCommentRepositiory(s),
	}
}
