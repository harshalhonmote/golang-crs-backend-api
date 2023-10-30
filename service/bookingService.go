package service

import (
	"github.com/harshalhonmote/crs/database"
	"github.com/harshalhonmote/crs/dto"
	"github.com/harshalhonmote/crs/models"
)

type Booking struct {
	database.Database
}

func (s *Booking) Create(bookingDto *dto.BookingAddDto) (string, models.Booking) {

	var car models.Car
	s.GetDB().Where("id =?", bookingDto.CarID).First(&car)

	var user models.User
	s.GetDB().Where("id =?", bookingDto.UserID).First(&user)

	var branch models.Branch
	s.GetDB().Where("id =?", bookingDto.BranchID).First(&branch)

	var booking models.Booking
	if car.Availability == true && user.KycStatus == true {
		booking = models.Booking{
			TripCost:      bookingDto.TripCost,
			BookingStatus: bookingDto.BookingStatus,
			BookingDate:   bookingDto.BookingDate,
			FromDate:      bookingDto.FromDate,
			ToDate:        bookingDto.ToDate,
			Branch:        &branch,
			User:          &user,
			Car:           &car,
		}

		resultBooking := s.GetDB().Create(&booking)
		if resultBooking.Error == nil {
			car.Availability = false
			s.GetDB().Save(&car)

			transaction := models.Transaction{
				Booking:     &booking,
				TransId:     bookingDto.TransId,
				Amount:      bookingDto.Amount,
				PaymentMode: bookingDto.PaymentMode,
			}

			resultTran := s.GetDB().Create(&transaction)
			if resultTran.Error != nil {
				return "Transaction error", booking
			} else {
				return "Booking success", booking
			}

		}
		return "Booking Error", booking
	}

	return "Car or user error", booking
}
func (s *Booking) GETBookingsOfUser(id int) []models.Booking {
	var bookings []models.Booking
	s.GetDB().Where("user_id =?", id).Preload("User").Preload("Branch").Preload("Car").Find(&bookings)
	return bookings
}
