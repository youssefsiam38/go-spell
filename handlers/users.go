package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefsiam38/spell/db"
	"github.com/youssefsiam38/spell/models"
	"github.com/youssefsiam38/spell/utils"
	"net/http"
)

type userWithToken struct {
	User  *models.User `json:"user"`
	Token *string      `json:"token"`
}

// Signup the user up
func Signup(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	
	retrievedUser, err := db.InsertUser(&user)

	if err != nil {
		// check if it is a duplication error
		if err.Error() == "Username is taken please pick another one" {
			utils.ErrResponse(c, 500, "Username is taken please pick another one")
			return
		}
	}

	token := utils.GenerateJWT(&user)

	UWT := userWithToken{
		User:  retrievedUser,
		Token: token,
	}

	c.JSON(http.StatusCreated, UWT)
}

// Login let the user to signin
func Login(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	retrievedUser, err := db.SelectUser(&user.Username)

	if err != nil {
		utils.ErrResponse(c, http.StatusNotFound, "Cannot find user with this credentials")
		return
	}
	if retrievedUser.Password != user.Password {
		utils.ErrResponse(c, http.StatusNotFound, "Incorrect password")
		return
	}

	token := utils.GenerateJWT(&user)

	UWT := userWithToken{
		User:  retrievedUser,
		Token: token,
	}

	c.JSON(http.StatusOK, UWT)

}

// GetMe is a handler to response with one user
func GetMe(c *gin.Context) {
	userPtr, _ := c.Get("userPtr")
	user := userPtr.(*models.User)

	c.JSON(200, user)
}

// GetUser is a handler to response with one user
func GetUser(c *gin.Context) {
	username := c.Param("username")

	user, err := db.SelectUser(&username)

	if err != nil {
		utils.ErrResponse(c, http.StatusNotFound, "User not found")
		return
	}

	c.JSON(200, struct{
		Username string `json:"username"`
		ID uint `json:"id"`
		Bio string `json:"bio"`
	}{
		Username: user.Username,
		ID: user.ID,
		Bio: user.Bio,
	})
}
