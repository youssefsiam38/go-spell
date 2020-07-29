package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/youssefsiam38/spell/models"
	"github.com/youssefsiam38/spell/db"
)


 // VerifyJWT token and return the user if verified
func VerifyJWT(tknStr string) (*models.User, error) {

	var claims Claims

	_ , err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	
	user, err := db.SelectUser(&claims.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}