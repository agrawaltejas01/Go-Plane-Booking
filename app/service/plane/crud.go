package PlaneService

import (
	planeRepo "Plane-Booking/app/repositories/planes"
	seatRepo "Plane-Booking/app/repositories/seats"
	models "Plane-Booking/db/Models"
)

func AddPlane(plane *models.Plane) (*models.Plane, error) {
	planeObject, err := planeRepo.AddPlane(plane)

	seatErr := seatRepo.AddSeats(int(planeObject.ID))

	if err != nil {
		return planeObject, err
	} else {
		return planeObject, seatErr
	}
}

func GetPlane(id string) (models.Plane, error) {
	return planeRepo.GetPlane(id)

}
