package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/SAHIL-Sharma21/go-taskForge/internal/model/todo"
	"github.com/SAHIL-Sharma21/go-taskForge/internal/repository"
	testing_pkg "github.com/SAHIL-Sharma21/go-taskForge/internal/testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoRepository_CreateTodo(t *testing.T) {
	_, testServer, cleanup := testing_pkg.SetupTest(t)
	defer cleanup()

	ctx := context.Background()
	todoRepo := repository.NewTodoRepository(testServer)

	t.Run("create todo successfully", func(t *testing.T) {
		userID := uuid.New().String()
		dueDate := time.Now().Add(24 * time.Hour)
		payload := &todo.CreateTodoPayload{
			Title:       "Test Todo",
			Description: testing_pkg.Ptr("Test todo description"),
			Priority:    testing_pkg.Ptr(todo.PriorityHigh),
			DueDate:     &dueDate,
		}

		result, err := todoRepo.CreateTodo(ctx, userID, payload)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.NotEqual(t, uuid.Nil, result.ID)
		assert.Equal(t, userID, result.UserID)
		assert.Equal(t, payload.Title, result.Title)
		assert.Equal(t, payload.Description, result.Description)
		assert.Equal(t, *payload.Priority, result.Priority)
		assert.Equal(t, payload.DueDate.Unix(), result.DueDate.Unix())
		assert.Equal(t, todo.StatusDraft, result.Status)
		assert.Nil(t, result.CompletedAt)
		testing_pkg.AssertTimestampsValid(t, result)
	})

	t.Run("create todo with minimum required fields", func(t *testing.T) {
		userID := uuid.New().String()
		payload := &todo.CreateTodoPayload{
			Title: "Minimal Todo",
		}

		result, err := todoRepo.CreateTodo(ctx, userID, payload)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, payload.Title, result.Title)
		assert.Nil(t, result.Description)
		assert.Equal(t, todo.PriorityMedium, result.Priority)
		assert.Nil(t, result.DueDate)
	})

	t.Run("create todo with metadata", func(t *testing.T) {
		userID := uuid.New().String()
		metadata := &todo.MetaData{
			Tags:  []string{"work", "urgent"},
			Color: testing_pkg.Ptr("#ff0000"),
		}
		payload := &todo.CreateTodoPayload{
			Title:    "Todo with Metadata",
			Metadata: metadata,
		}

		result, err := todoRepo.CreateTodo(ctx, userID, payload)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, metadata.Tags, result.MetaData.Tags)
		assert.Equal(t, metadata.Color, result.MetaData.Color)
	})

	t.Run("with canceled context", func(t *testing.T) {
		canceledCtx, cancel := context.WithCancel(ctx)
		cancel()

		userID := uuid.New().String()
		payload := &todo.CreateTodoPayload{
			Title: "Canceled Todo",
		}

		result, err := todoRepo.CreateTodo(canceledCtx, userID, payload)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
