package UserService

import (
	userRepo "Plane-Booking/app/repositories/user"
	models "Plane-Booking/db/Models"
)

func AddUser(user *models.User) (*models.User, error) {
	userObject, err := userRepo.AddUser(user)

	return userObject, err
}

func GetUser(id string) (models.User, error) {
	return userRepo.GetUser(id)

}
