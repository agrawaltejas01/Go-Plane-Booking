package bookingRepo

import (
	"Plane-Booking/db"
	models "Plane-Booking/db/Models"
)

func Book(booking *models.Booking) (*models.Booking, error) {

	err := db.Database.Create(&booking).Error

	if err != nil {
		return &models.Booking{}, err
	}
	return booking, nil

}

func ShowBookings(id string) (models.Booking, error) {
	var booking models.Booking

	err := db.Database.Where("id = ?", id).Find(&booking).Error

	if err != nil {
		return models.Booking{}, err
	}

	return booking, err
}
