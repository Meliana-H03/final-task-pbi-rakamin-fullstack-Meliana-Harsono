package models

import (
	"time"
)

type Photos struct {
	Id        int       `gorm:"primaryKey;not null" json:"id"`
	Title     string    `gorm:"type:varchar(300);not null" json:"title"`
	Caption   string    `gorm:"type:varchar(300);not null" json:"caption"`
	PhotoURL  string    `gorm:"type:varchar(300);not null" json:"photo_url"`
	UserID    int       `gorm:"foreignKey;not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
