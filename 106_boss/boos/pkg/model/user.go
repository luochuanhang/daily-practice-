package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Age      int    `json:"age,omitempty"`
}
