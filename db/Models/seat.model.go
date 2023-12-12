package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model

	PlaneId int `gorm:"size:255;not null;foreignKey:Plane" json:"planeId"`
	No      int `gorm:"not null;" json:"no"`
}
