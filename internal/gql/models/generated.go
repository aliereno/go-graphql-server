// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"

	"github.com/aliereno/go-graphql-server/internal/orm/models"
)

type LoginModel struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type OrderByTypes string

const (
	OrderByTypesCreatedAtAsc  OrderByTypes = "createdAt_ASC"
	OrderByTypesCreatedAtDesc OrderByTypes = "createdAt_DESC"
	OrderByTypesIDDesc        OrderByTypes = "id_DESC"
	OrderByTypesIDAsc         OrderByTypes = "id_ASC"
)

var AllOrderByTypes = []OrderByTypes{
	OrderByTypesCreatedAtAsc,
	OrderByTypesCreatedAtDesc,
	OrderByTypesIDDesc,
	OrderByTypesIDAsc,
}

func (e OrderByTypes) IsValid() bool {
	switch e {
	case OrderByTypesCreatedAtAsc, OrderByTypesCreatedAtDesc, OrderByTypesIDDesc, OrderByTypesIDAsc:
		return true
	}
	return false
}

func (e OrderByTypes) String() string {
	return string(e)
}

func (e *OrderByTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderByTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid orderByTypes", str)
	}
	return nil
}

func (e OrderByTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
