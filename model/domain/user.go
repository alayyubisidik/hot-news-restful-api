package domain

import "time"

type User struct {
	ID        int `gorm:"primary_key;autoIncrement"`
	Username  string
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt  time.Time
	Article   []Article `gorm:"foreignKey:user_id;references:id"`
}
