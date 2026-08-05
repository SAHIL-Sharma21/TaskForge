package repository

import "github.com/SAHIL-Sharma21/go-taskForge/internal/server"

type Repositories struct {
	Todo     *TodoRepository
	Comment  *CommentRepository
	Category *CategoryRepository
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{
		Todo:     NewTodoRepository(s),
		Comment:  NewCommentRepositiory(s),
		Category: NewCategoryRepository(s),
	}
}
