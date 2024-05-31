package models

import (
	"time"
)

type Register struct {
	Id        int       `gorm:"primaryKey;not null" json:"id"`
	Username  string    `gorm:"type:varchar(300);not null" json:"username"`
	Email     string    `gorm:"type:varchar(300);not null; unique" json:"email"`
	Password  string    `gorm:"type:varchar(300);not null" validate:"min=6" json:"password"`
	Photos    []Photos  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
