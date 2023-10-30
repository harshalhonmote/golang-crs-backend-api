package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/controllers"
)

func BookingRoutes(router *gin.Engine) {
	booking := new(controllers.Booking)
	router.POST("/booking/add", booking.AddBooking)
	router.GET("/booking/:id", booking.GETBookingsOfUser)
	// router.POST("/user/signin", user.Login)
	// router.POST("/user/signup", user.RegisterUser)
	// router.POST("/user/update", user.Update)
}
