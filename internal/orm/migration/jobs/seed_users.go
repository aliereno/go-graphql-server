package jobs

import (
	"fmt"
	"github.com/aliereno/go-graphql-server/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		name := ""
		email := ""
		for i := 1; i <= 5; i++ {
			name = fmt.Sprintf("User%d", i)
			email = fmt.Sprintf("user%d@gmail.com", i)
			dbo := &models.User{
				Name:  &name,
				Email: email,
			}
			// Create scoped clean db interface
			tx := db.New().Begin()
			tx = tx.Create(dbo).First(dbo)

			tx = tx.Commit()
		}
		return nil
	},
	Rollback: func(db *gorm.DB) error {
		return nil
	},
}
var SeedBooks *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_BOOKS",
	Migrate: func(db *gorm.DB) error {
		name := ""
		userID := 0

		for i := 1; i <= 5; i++ {
			name = fmt.Sprintf("Book%d", i)
			userID = i
			dbo := &models.Book{
				Name:   name,
				UserID: &userID,
			}
			// Create scoped clean db interface
			tx := db.New().Begin()
			tx = tx.Create(dbo).First(dbo)

			tx = tx.Commit()
		}
		return nil
	},
	Rollback: func(db *gorm.DB) error {
		return nil
	},
}
