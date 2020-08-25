package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/taroodr/til/go/graphql-model-relation-test/domain"
	"github.com/taroodr/til/go/graphql-model-relation-test/graph/generated"
	"github.com/taroodr/til/go/graphql-model-relation-test/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	user, err := r.UserRepository.Create(ctx, &domain.User{
		Name: input.Name,
	})

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID: string(user.ID),
		Name: user.Name,
	}, nil

}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	log.Printf("[queryResolver.User] id: %s", id)
	user, err := r.UserRepository.Get(ctx, domain.UserID(id))
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID: string(user.ID),
		Name: user.Name,
	}, nil

}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
