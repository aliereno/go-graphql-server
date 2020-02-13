package resolvers

import (
	"context"
	"github.com/aliereno/go-graphql-server/internal/helpers"

	//"github.com/aliereno/go-graphql-server/internal/logger"
	"strconv"

	gqlModels "github.com/aliereno/go-graphql-server/internal/gql/models"
	"github.com/aliereno/go-graphql-server/internal/orm/models"
)

func (r *mutationResolver) CreateBook(ctx context.Context, name *string, description *string, userID *string) (*models.Book, error) {
	userIDToInt64, _ := strconv.ParseInt(*userID, 10, 32)
	userIDToInt := int(userIDToInt64)
	dbo := &models.Book{
		Name:        *name,
		Description: description,
		UserID:      &userIDToInt,
	}
	db := r.ORM.DB.New().Begin()
	db = db.Create(dbo).First(dbo)
	db = db.Commit()
	return dbo, db.Error
}
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, name *string, description *string, userID *string) (*models.Book, error) {
	userIDToInt64, _ := strconv.ParseInt(*userID, 10, 32)
	userIDToInt := int(userIDToInt64)
	dbo := &models.Book{
		Name:        *name,
		Description: description,
		UserID:      &userIDToInt,
	}
	model := &models.Book{}
	tx := r.ORM.DB.New().Begin()
	tx = tx.Where("id = ?", id).First(&model).Update(dbo).First(dbo)
	tx = tx.Commit()
	return dbo, tx.Error
}
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*models.Book, error) {
	book := models.Book{}
	tx := r.ORM.DB.New().Begin()
	tx = tx.Where("id = ?", id).First(&book).Delete(&book)
	tx = tx.Commit()
	return &book, tx.Error
}
func (r *queryResolver) Books(ctx context.Context, filter map[string]interface{}, orderBy *gqlModels.OrderByTypes, limit *int) ([]*models.Book, error) {
	var dbRecords []*models.Book
	db := r.ORM.DB.New()
	db = db.Preload("User").Where(filter).Order(helpers.OrderToString(orderBy)).Find(&dbRecords)
	return dbRecords, db.Error
}
