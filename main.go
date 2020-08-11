package main

import (
	// dev
	// "github.com/gin-contrib/pprof"
	"github.com/gin-contrib/cors"
	_ "github.com/joho/godotenv/autoload"
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

	// dev
	// pprof.Register(r)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3001", "http://127.0.0.1:3000", "http://localhost:3001", "http://localhost:3000", "http://localhost"}, ///// env
		AllowMethods:     []string{"POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
