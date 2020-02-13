package orm

var autoMigrate, logMode bool
var hostDB, userDB, passwordDB, nameDB, sslDB, dialect string

func init() {
	dialect = ""
	hostDB = ""
	userDB = ""
	passwordDB = ""
	nameDB = ""
	sslDB = "disable"
	logMode = true
	autoMigrate = true
}
