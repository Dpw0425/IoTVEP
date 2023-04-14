package entity

// 用户实体
type User struct {
	ID       uint   `gorm:"primaryKey" binding:"omitempty"`
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty"`
}
