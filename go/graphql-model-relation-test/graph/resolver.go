package graph

import "github.com/taroodr/til/go/graphql-model-relation-test/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	TodoRepository domain.TodoRepository
	UserRepository domain.UserRepository
}

