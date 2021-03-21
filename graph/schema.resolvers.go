package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"

	"github.com/eivy/aptitude_bulb/db"
	"github.com/eivy/aptitude_bulb/graph/generated"
	"github.com/eivy/aptitude_bulb/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	u, err := con.CreateUser(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   int(u.ID),
		Name: u.Name,
	}, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	item, err := con.CreateItem(ctx, db.CreateItemParams{
		Name:      input.Name,
		Location:  sql.NullString{String: input.Location},
		Counts:    int32(input.Counts),
		ManagerID: int32(input.Manager),
	})
	if err != nil {
		return nil, err
	}
	u, err := con.GetUser(ctx, item.ManagerID)
	if err != nil {
		return nil, err
	}
	val := &model.Item{
		ID:     int(item.ID),
		Name:   item.Name,
		Counts: int(item.Counts),
		Manager: &model.User{
			ID:   int(u.ID),
			Name: u.Name,
		},
	}
	if item.Location.Valid {
		val.Location = &item.Location.String
	} else {
		val.Location = nil
	}
	return val, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.User, error) {
	u, err := con.DeleteUser(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   int(u.ID),
		Name: u.Name,
	}, nil
}

func (r *mutationResolver) DeleteItem(ctx context.Context, id int) (*model.Item, error) {
	item, err := con.DeleteItem(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	u, err := con.GetUser(ctx, item.ManagerID)
	if err != nil {
		return nil, err
	}
	val := &model.Item{
		ID:     int(item.ID),
		Name:   item.Name,
		Counts: int(item.Counts),
		Manager: &model.User{
			ID:   int(u.ID),
			Name: u.Name,
		},
	}
	if item.Location.Valid {
		val.Location = &item.Location.String
	} else {
		val.Location = nil
	}
	return val, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	u, err := con.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, len(u))
	for i, v := range u {
		users[i] = &model.User{
			ID:   int(v.ID),
			Name: v.Name,
		}
	}
	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	u, err := con.GetUser(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   int(u.ID),
		Name: u.Name,
	}, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	item, err := con.ListItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*model.Item, len(item))
	for i, v := range item {
		items[i] = &model.Item{
			ID:   int(v.ID),
			Name: v.Name,
			Manager: &model.User{
				ID:   int(v.ID_2),
				Name: v.Name_2,
			},
		}
		if v.Location.Valid {
			items[i].Location = &v.Location.String
		} else {
			items[i].Location = nil
		}
	}
	return items, nil
}

func (r *queryResolver) Item(ctx context.Context, id int) (*model.Item, error) {
	item, err := con.GetItem(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	val := &model.Item{
		ID:     int(item.ID),
		Name:   item.Name,
		Counts: int(item.Counts),
		Manager: &model.User{
			ID:   int(item.ID_2),
			Name: item.Name_2,
		},
	}
	if item.Location.Valid {
		val.Location = &item.Location.String
	} else {
		val.Location = nil
	}
	return val, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
