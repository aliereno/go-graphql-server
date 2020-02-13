package server

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = "localhost"
	port = "7778"
	gqlPath = "/graphql"
	gqlPgPath = "/"
	isPgEnabled = true
}
