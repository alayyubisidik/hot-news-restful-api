package domain

import "time"

type Comment struct {
	ID        int `gorm:"primary_key;autoIncrement"`
	ArticleId int
	UserId    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User    `gorm:"foreignKey:user_id;references:id"`
	Article   Article `gorm:"foreignKey:article_id;references:id"`
}
