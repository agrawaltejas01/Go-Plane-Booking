package bookingService

import (
	bookingRepo "Plane-Booking/app/repositories/booking"
	models "Plane-Booking/db/Models"
)

func BookSeat(booking *models.Booking) (*models.Booking, error) {
	return bookingRepo.Book(booking)
}

func FindBooking(bookingId string) (models.Booking, error) {
	return bookingRepo.ShowBookings(bookingId)
}
