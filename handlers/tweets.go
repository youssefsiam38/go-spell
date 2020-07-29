package handlers

import (
	"github.com/youssefsiam38/spell/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/youssefsiam38/spell/db"
	"github.com/youssefsiam38/spell/models"
	"strconv"
)

// Tweet is handler to insert a user's tweet in the database
func Tweet(c *gin.Context) {
	userPtr, _ := c.Get("userPtr")
	user := userPtr.(*models.User)

	var tweet models.Tweet

	c.BindJSON(&tweet)

	tweet.UserID = user.ID

	err := db.InsertTweet(&tweet)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, tweet)
}

//GetTweets get all tweets for the user in param
func GetTweets(c *gin.Context) {

	username := c.Param("username")

	user, err := db.SelectUser(&username)

	if err != nil {
		utils.ErrResponse(c, http.StatusNotFound, "User not found")
		return
	}

	tweets, err := db.SelectTweets(user)

	if err != nil {
		panic(err)
	}

	c.JSON(200, tweets)
}

// DeleteTweet is a handler to delete a tweet
func DeleteTweet(c *gin.Context) {
	userPtr, _ := c.Get("userPtr")
	user := userPtr.(*models.User)

	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}

	id := uint(id64)

	tweet, err := db.SelectTweet(&id)

	if err != nil {
		if err.Error() != "tweet not found" {
			panic(err)
		} else {
			utils.ErrResponse(c, http.StatusNotFound, err.Error())
			return
		}
	}

	if tweet.UserID != user.ID {
		utils.ErrResponse(c, http.StatusForbidden, "Cannot delete this tweet")
		return
	}

	if err = db.DeleteTweet(&id); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, tweet)
}
