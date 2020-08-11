package db

import (
	"database/sql"
	"os"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// CreateDBIfNotExists Creates DB If Not Exists :D
func CreateDBIfNotExists() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+`:`+os.Getenv("MYSQL_PASS")+`@tcp(`+os.Getenv("MYSQL_HOST")+`:`+os.Getenv("MYSQL_PORT")+`)/`)
	defer db.Close()

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("create database if not exists spell")
	if err != nil {
		panic(err)
	}
	Connect()
	createUsersTable()
	createTweetsTable()

}

// Connect to database
func Connect() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+`:`+os.Getenv("MYSQL_PASS")+`@tcp(`+os.Getenv("MYSQL_HOST")+`:`+os.Getenv("MYSQL_PORT")+`)/spell`)
	if err != nil {
		panic(err)
	}
	duration, err := time.ParseDuration("40ms")
	defer db.SetConnMaxLifetime(duration)

	return db
}

func createUsersTable() error {
	db := Connect()
	defer db.Close()

	_, err := db.Exec(`
		create table if not exists users
		(username varchar(50) not null UNIQUE ,
		id int AUTO_INCREMENT UNIQUE not null,
		password varchar(50) not null,
		bio varchar(225),
		primary key (id));
	`)
	if err != nil {
		panic(err)
	}

	return nil
}

func createTweetsTable() error {
	db := Connect()
	defer db.Close()


	_, err := db.Exec(`
		create table if not exists tweets (
		id int AUTO_INCREMENT,
		body varchar(225) not null,
		userID int not null,
		createdAT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		primary key (id),
		foreign key (userID) references users(id) on delete cascade);
	`)
	if err != nil {
		panic(err)
	}

	return nil
}
