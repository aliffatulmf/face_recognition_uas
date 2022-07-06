package model

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID         uint           `gorm:"primaryKey"`
	Name       string         ``
	Image      string         ``
	Identifier string         ``
	CreatedAt  time.Time      ``
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
