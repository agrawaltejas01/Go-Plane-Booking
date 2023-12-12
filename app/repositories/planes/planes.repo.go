package userRepo

import (
	"Plane-Booking/db"
	models "Plane-Booking/db/Models"
)

func AddPlane(plane *models.Plane) (*models.Plane, error) {

	err := db.Database.Create(&plane).Error

	if err != nil {
		return &models.Plane{}, err
	}
	return plane, nil

}

func GetPlane(id string) (models.Plane, error) {
	var plane models.Plane

	err := db.Database.Where("id = ?", id).Find(&plane).Error

	if err != nil {
		return models.Plane{}, err
	}

	return plane, err
}
