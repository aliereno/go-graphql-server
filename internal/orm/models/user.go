package models

// User defines a user for the app
type User struct {
	BaseModelSoftDelete        // We don't to actually delete the users, maybe audit if we want to hard delete them? or wait x days to purge from the table, also
	Email               string `gorm:"not null;unique_index:idx_email"`
	Name                *string
	Books               []*Book `gorm:"foreignkey:UserID"`
}
