package models

// Tweet of a user
type Tweet struct {
	ID        uint   `json:"id"`
	Body      string `json:"body"`
	UserID    uint   `json:"userID"`
	CreatedAt string `json:"createdAt"`
}
