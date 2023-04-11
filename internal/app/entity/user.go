package entity

type User struct {
	ID       int    `gorm:"primaryKey" binding:"omitempty"`
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty"`
}

type LoginEmail struct {
	Email        string `json:"email"`
	Verification string `json:"verification"`
}
