// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	"fmt"
	log "github.com/aliereno/go-graphql-server/internal/logger"
	"github.com/aliereno/go-graphql-server/internal/orm/migration"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", hostDB, userDB, nameDB, sslDB, passwordDB) //Build connection string

	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(logMode)
	// Automigrate tables
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Info("[ORM] Database connection initialized.")
	return orm, err
}
