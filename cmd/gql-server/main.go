package main

import (
	log "github.com/aliereno/go-graphql-server/internal/logger"

	"github.com/aliereno/go-graphql-server/internal/orm"
	"github.com/aliereno/go-graphql-server/pkg/server"
)

func main() {
	// Create a new ORM instance to send it to our
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	// Send: ORM instance
	server.Run(orm)
}
