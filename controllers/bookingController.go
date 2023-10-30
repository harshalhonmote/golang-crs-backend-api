package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/dto"
	"github.com/harshalhonmote/crs/helper"
	"github.com/harshalhonmote/crs/service"
)

type Booking struct {
	helper.Controller
}

func (u *Booking) AddBooking(c *gin.Context) {
	var params *dto.BookingAddDto
	err := c.BindJSON(&params)
	fmt.Println(err)
	fmt.Println(params.CarID, ":", params.BookingDate)

	result, _ := new(service.Booking).Create(params)
	c.JSON(http.StatusOK, result)
	//u.Json(c, result, data)
}

//json Api for AddBooking
/*
{
	"tripCost":200,
	"branchID":1,
	"bookingStatus":"scheduled",
	"userID":1,
	"carID":1,
	"bookingDate":"2023-03-08T15:04:05Z",
  "fromDate":"2023-03-08T15:04:05Z",
  "toDate":"2023-03-08T15:04:05Z"

}
*/
func (u *Booking) GETBookingsOfUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	bookings := new(service.Booking).GETBookingsOfUser(userID)
	c.JSON(http.StatusOK, bookings)
}
