package SeatService

import (
	planeRepo "Plane-Booking/app/repositories/planes"
	models "Plane-Booking/db/Models"
)

func AddPlane(plane *models.Plane) (*models.Plane, error) {
	planeObject, err := planeRepo.AddPlane(plane)

	return planeObject, err
}
