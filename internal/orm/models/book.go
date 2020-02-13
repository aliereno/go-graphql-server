package models

// User defines a user for the app
type Book struct {
	BaseModelSoftDelete // We don't to actually delete the users, maybe audit if we want to hard delete them? or wait x days to purge from the table, also
	Name                string
	Description         *string `gorm:"size:1048"`
	UserID              *int
	User                *User
}
