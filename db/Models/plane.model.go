package models

import "gorm.io/gorm"

type Plane struct {
	gorm.Model

	Name        string `gorm:"size:255;not null;" json:"name"`
	Source      string `gorm:"size:255;not null;" json:"source"`
	Destination string `gorm:"size:255;not null;" json:"destination"`
}
