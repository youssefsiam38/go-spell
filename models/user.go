package models

// User in the app
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Password string `json:"password"`
}
