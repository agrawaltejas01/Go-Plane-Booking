package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model

	PlaneId int   `gorm:"not null;" json:"planeId"`
	Plane   Plane `gorm:"foreignKey:PlaneId"`
	SeatId  int   `gorm:"not null;" json:"seatId"`
	Seat    Seat  `gorm:"not null;" json:"Seat"`
	UserId  int   `gorm:"not null;" json:"userId"`
	User    User  `gorm:"not null;" json:"User"`
}
