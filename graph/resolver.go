package graph

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/amitbikram/gql-go/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	va, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		panic("pamicked")
	}
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", va),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
