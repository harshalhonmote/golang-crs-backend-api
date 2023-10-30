package service

import (
	"fmt"
	"time"

	"github.com/harshalhonmote/crs/database"
	"github.com/harshalhonmote/crs/dto"
	"github.com/harshalhonmote/crs/models"
	"gorm.io/gorm"
)

type Car struct {
	database.Database
}

func (s *Car) Create(car *models.Car) (*gorm.DB, *models.Car) {

	var branch models.Branch
	//s.GetDB().Find(&branch, "id=?", car.BranchID)
	fmt.Println(car.BranchID)
	s.GetDB().Where("id =?", car.BranchID).First(&branch)
	fmt.Println("data", branch)

	carTemp := &models.Car{
		Brand:        car.Brand,
		Model:        car.Model,
		Availability: car.Availability,
		ImageUrl:     car.ImageUrl,
		CreatedAt:    time.Now(),
		Branch:       &branch,
	}
	fmt.Println(carTemp)
	result := s.GetDB().Create(&carTemp)
	return result, car
}

func (s *Car) GetAllCars() (*gorm.DB, []models.Car) {
	var car []models.Car
	result := s.GetDB().Find(&car)
	return result, car
}
func (s *Car) GetCar(id int) (*gorm.DB, *models.Car) {
	var car models.Car
	result := s.GetDB().Where("id =?", id).Preload("Branch").First(&car)
	return result, &car
}

func (s *Car) ChangeAvailability(carDto *dto.CarAvailableDto) (*gorm.DB, *models.Car) {
	var car models.Car
	s.GetDB().Where("id =?", carDto.Id).First(&car)
	car.Availability = carDto.Availability
	result := s.GetDB().Save(&car)
	return result, &car
}
