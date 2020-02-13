package auth

import (
	"context"
	"github.com/aliereno/go-graphql-server/internal/orm/models"
	"github.com/aliereno/go-graphql-server/pkg/utils"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"time"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

// Ping is simple keep-alive/ping handler
func GetToken(user models.User) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 30).Unix(), // TOKEN EXPIRE TIME
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mysupersecretpassword))
	if err != nil {
		return "Could not generate token", err
	}
	return tokenString, nil
}

func LookToken(c context.Context) error {
	gc, err1 := utils.GinContextFromContext(c)
	if err1 != nil {
		return err1
	}
	_, err := request.ParseFromRequest(gc.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
		b := []byte(mysupersecretpassword)
		return b, nil
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}
