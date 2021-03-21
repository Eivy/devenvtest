package graph

import "github.com/eivy/aptitude_bulb/graph/model"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Users []model.User
	Items []model.Item
	User  *model.User
	Item  *model.Item
}
