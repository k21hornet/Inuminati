package model

import "time"

type Dog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Img       string    `json:"img" gorm:"not null"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type DogResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Img       string    `json:"img" gorm:"not null"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}
