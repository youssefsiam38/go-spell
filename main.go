package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/gin-contrib/cors"
	"github.com/youssefsiam38/spell/db"
	// "github.com/youssefsiam38/spell/models"
	// "github.com/youssefsiam38/spell/utils"
	// "net/http"
	// "time"
	"os"
	// "fmt"
	"github.com/youssefsiam38/spell/handlers"
	"github.com/youssefsiam38/spell/middlewares"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {

	db.CreateDBIfNotExists()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3001", "http://127.0.0.1:3000", "http://localhost:3001", "http://localhost:3000"}, ///// env
		AllowMethods:     []string{"POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))

	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	//////////////////
	// user handlers
	//////////////////

	r.POST("/signup", handlers.Signup)

	r.POST("/login", handlers.Login)
	
	// get the authintecated user info
	r.GET("/me", middlewares.Auth, handlers.GetMe)

	// get a user's info
	r.GET("/user/:username", handlers.GetUser)

	//////////////////
	// tweets handlers
	//////////////////

	// post tweet
	r.POST("/tweet", middlewares.Auth, handlers.Tweet)
	
	// get user's tweets
	r.GET("/user/:username/tweets", handlers.GetTweets)

	// delete tweet
	r.DELETE("/delete/tweet/:id", middlewares.Auth, handlers.DeleteTweet)

	r.Run(`:` + os.Getenv("PORT"))

}