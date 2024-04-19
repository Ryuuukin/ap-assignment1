package models

import "time"

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
	Game      string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
