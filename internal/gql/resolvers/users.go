package resolvers

import (
	"context"
	gqlModels "github.com/aliereno/go-graphql-server/internal/gql/models"
	"github.com/aliereno/go-graphql-server/internal/handlers/auth"
	"github.com/aliereno/go-graphql-server/internal/helpers"

	"github.com/aliereno/go-graphql-server/internal/orm/models"
)

func (r *mutationResolver) LoginUser(ctx context.Context, email string, password *string) (*gqlModels.LoginModel, error) {
	record := &gqlModels.LoginModel{}
	user := models.User{}
	db := r.ORM.DB.New()
	db = db.Where("email = ?", email).First(&user)

	record.User = &user
	token, err := auth.GetToken(user)
	record.Token = token

	return record, err
}
func (r *mutationResolver) CreateUser(ctx context.Context, name *string, email *string) (*models.User, error) {
	dbo := &models.User{
		Name:  name,
		Email: *email,
	}
	tx := r.ORM.DB.New().Begin()
	tx = tx.Create(dbo).First(dbo)
	tx = tx.Commit()
	return dbo, tx.Error
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name *string, email *string) (*models.User, error) {
	dbo := &models.User{
		Name:  name,
		Email: *email,
	}
	model := &models.User{}
	tx := r.ORM.DB.New().Begin()
	tx = tx.Where("id = ?", id).First(&model).Update(dbo).First(dbo)
	tx = tx.Commit()
	return dbo, tx.Error
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	user := models.User{}
	tx := r.ORM.DB.New().Begin()
	tx = tx.Where("id = ?", id).First(&user).Delete(&user)
	tx = tx.Commit()
	return &user, tx.Error
}
func (r *queryResolver) Users(ctx context.Context, filter map[string]interface{}, orderBy *gqlModels.OrderByTypes, limit *int) ([]*models.User, error) {
	var dbRecords []*models.User
	db := r.ORM.DB.New()
	db = db.Preload("Books").Where(filter).Order(helpers.OrderToString(orderBy)).Find(&dbRecords)
	return dbRecords, db.Error
}
