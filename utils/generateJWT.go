package utils

import (
	// import the jwt-go library
	"github.com/dgrijalva/jwt-go"
	"github.com/youssefsiam38/spell/models"
	"os"
)

// JWT key used to create the signature
var jwtKey = []byte(os.Getenv("jwtKey"))

// Claims is the credentials to be stored in the token
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Valid is just to make Claims implement jwt.Claims
func (c Claims) Valid() error {
	return nil
}

// GenerateJWT token
func GenerateJWT(userptr *models.User) *string {
	user := *userptr
	claims := Claims{
		Username: user.Username,
		Password: user.Password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		panic(err)
	}
	return &tokenString
}
