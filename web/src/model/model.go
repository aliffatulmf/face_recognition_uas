package model

import (
	"time"

	"gorm.io/gorm"
)

type Dataset struct {
	ID         uint           `gorm:"primaryKey"`
	Identifier string         `form:"identifier" binding:"required" gorm:"index:idx_identifier,unique"`
	Name       string         `form:"name" binding:"required"`
	CreatedAt  time.Time      ``
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
