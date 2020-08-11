package db

import (
	"errors"
	"fmt"
	"github.com/youssefsiam38/spell/models"
	"strconv"
	"strings"
)

// InsertUser inserts a user to the database
func InsertUser(user *models.User) (*models.User, error) {
	db := Connect()
	defer db.Close()

	result, err := db.Exec(`
		insert into users (username, password, bio) 
		values (?, ?, ?);
	`, user.Username, user.Password, user.Bio)
	if err != nil {

		// mysql error status code
		code, _ := strconv.ParseInt(strings.Replace(strings.Fields(err.Error())[1], ":", "", 1), 10, 16)

		if code == 1062 {
			return nil, errors.New("Username is taken please pick another one")
		}

		return nil, err
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsEffected != 1 {
		return nil, errors.New("Somtheing went wrong with inserting in database")
	}

	user, _ = SelectUser(&user.Username)
	return user, nil
}

// SelectUser gets a user from the database using the username
func SelectUser(username *string) (*models.User, error) {

	db := Connect()
	defer db.Close()

	var user models.User

	err := db.QueryRow(`
		select id, username, password, bio from users where username = ?
	`, *username).Scan(&user.ID, &user.Username, &user.Password, &user.Bio)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// InsertTweet is a function to insert a tweet to the database
func InsertTweet(tweet *models.Tweet) error {
	db := Connect()
	defer db.Close()

	_, err := db.Exec(`
		insert into tweets (body, userID)
		values (?, ?);
	`, tweet.Body, tweet.UserID)

	if err != nil {
		return err
	}

	db.QueryRow(`
		select id, createdAT from tweets where userID= ? order by createdAT desc;
	`, tweet.UserID).Scan(&tweet.ID, &tweet.CreatedAt)

	return nil
}

// SelectTweet is to select a single tweet from the database
func SelectTweet(id *uint) (*models.Tweet, error) {
	db := Connect()
	defer db.Close()

	var tweet models.Tweet

	row := db.QueryRow(`
	select * from tweets where id= ?;
	`, *id)

	err := row.Scan(&tweet.ID, &tweet.Body, &tweet.UserID, &tweet.CreatedAt)
	if err != nil || &tweet == nil {
		return nil, errors.New("tweet not found")
	}

	fmt.Println(tweet)

	return &tweet, nil
}

// SelectTweets get all tweets in the database tweeted by this user
func SelectTweets(user *models.User) ([]models.Tweet, error) {
	db := Connect()
	defer db.Close()

	tweets := []models.Tweet{}

	rows, err := db.Query(`
		select id, body, CreatedAT from tweets where userID= ? order by createdAT desc;
	`, user.ID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tweet models.Tweet

		rows.Scan(&tweet.ID, &tweet.Body, &tweet.CreatedAt)

		tweet.UserID = user.ID

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}

// DeleteTweet sends query to delete specific tweet
func DeleteTweet(id *uint) error {
	db := Connect()
	defer db.Close()

	result, err := db.Exec(`
	delete from tweets where id= ?;
	`, *id)

	fmt.Println(result.RowsAffected())
	return err
}
