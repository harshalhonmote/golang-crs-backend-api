package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/dto"
	"github.com/harshalhonmote/crs/helper"
	"github.com/harshalhonmote/crs/models"
	"github.com/harshalhonmote/crs/service"
)

type Car struct {
	helper.Controller
}

func (car *Car) AddCar(c *gin.Context) {
	var params *models.Car
	c.BindJSON(&params)
	result, data := new(service.Car).Create(params)
	//fmt.Println("controller...", params.BranchID)

	car.Json(c, result, data)
}

func (car *Car) GetAllCars(c *gin.Context) {
	result, data := new(service.Car).GetAllCars()
	car.Json(c, result, data)
}
func (car *Car) GetCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, data := new(service.Car).GetCar(id)
	car.Json(c, result, data)
}
func (car *Car) ChangeAvailability(c *gin.Context) {
	var carDto *dto.CarAvailableDto
	c.BindJSON(&carDto)
	result, data := new(service.Car).ChangeAvailability(carDto)
	car.Json(c, result, data)
}
