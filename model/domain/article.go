package domain

import "time"

type Article struct {
	ID         int `gorm:"primary_key;autoIncrement"`
	Title      string
	Slug       string
	Content    string
	UserId     int	`gorm:"column:user_id"`
	CategoryId int `gorm:"column:category_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User `gorm:"foreignKey:user_id;references:id"`
	Category   Category `gorm:"foreignKey:category_id;references:id"`
}
