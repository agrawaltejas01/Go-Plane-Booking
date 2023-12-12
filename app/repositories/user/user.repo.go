package userRepo

import (
	"Plane-Booking/db"
	models "Plane-Booking/db/Models"
)

func AddUser(user *models.User) (*models.User, error) {

	err := db.Database.Create(&user).Error

	if err != nil {
		return &models.User{}, err
	}
	return user, nil

}

func GetUser(id string) (models.User, error) {
	var user models.User

	err := db.Database.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, err
}
