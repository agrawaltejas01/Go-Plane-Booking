package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name string `gorm:"size:255;not null;" json:"name"`
}
