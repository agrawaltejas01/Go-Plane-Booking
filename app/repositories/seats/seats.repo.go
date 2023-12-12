package userRepo

import (
	"Plane-Booking/db"
	models "Plane-Booking/db/Models"
)

func AddSeats(planeId int) error {

	seats := []*models.Seat{}

	for i := 1; i <= 10; i++ {

		currSeat := models.Seat{
			No:      i,
			PlaneId: planeId,
		}

		seats = append(seats, &currSeat)
	}

	err := db.Database.Create(seats).Error

	return err

}
