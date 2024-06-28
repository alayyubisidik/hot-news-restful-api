package domain

import "time"

type Category struct {
	ID        int `gorm:"primary_key;autoIncrement"`
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Article   []Article `gorm:"foreignKey:user_id;references:id"`
}
