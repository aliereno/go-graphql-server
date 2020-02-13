package resolvers

import (
	"context"

	"github.com/aliereno/go-graphql-server/internal/gql"
	"github.com/aliereno/go-graphql-server/internal/gql/models"
	models1 "github.com/aliereno/go-graphql-server/internal/orm/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) LoginUser(ctx context.Context, email string, password *string) (*models.LoginModel, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, name *string, email *string) (*models1.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name *string, email *string) (*models1.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models1.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateBook(ctx context.Context, name *string, description *string, userID *string) (*models1.Book, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, name *string, description *string, userID *string) (*models1.Book, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*models1.Book, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, filter map[string]interface{}, orderBy *models.OrderByTypes, limit *int) ([]*models1.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Books(ctx context.Context, filter map[string]interface{}, orderBy *models.OrderByTypes, limit *int) ([]*models1.Book, error) {
	panic("not implemented")
}
