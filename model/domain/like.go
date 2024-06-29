package domain

import "time"

type Like struct {
	ID        int `gorm:"primary_key;autoIncrement"`
	UserId    int
	ArticleId int
	CreatedAt time.Time
	User      User     `gorm:"foreignKey:user_id;references:id"`
	Article  Article `gorm:"foreignKey:article_id;references:id"`
}
