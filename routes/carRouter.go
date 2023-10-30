package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/controllers"
)

func CarRoutes(router *gin.Engine) {
	car := new(controllers.Car)
	router.GET("/car/list", car.GetAllCars)
	router.POST("/car/add", car.AddCar)
	router.GET("/car/get/:id", car.GetCar)

	// router.POST("/user/signup", car.RegisterUser)
	// router.POST("/user/update", car.Update)
}
