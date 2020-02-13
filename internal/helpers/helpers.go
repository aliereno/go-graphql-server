package helpers

import "github.com/aliereno/go-graphql-server/internal/gql/models"

func OrderToString(orderBy *models.OrderByTypes) string {
	if orderBy != nil {
		var dbOrder string
		if orderBy.String() == "createdAt_ASC" {
			dbOrder = "created_at asc"
		} else if orderBy.String() == "createdAt_DESC" {
			dbOrder = "created_at desc"
		} else if orderBy.String() == "id_ASC" {
			dbOrder = "id asc"
		} else if orderBy.String() == "id_DESC" {
			dbOrder = "id desc"
		}
		return dbOrder
	}
	return "id asc"
}
