package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/controllers"
)

func UserRoutes(router *gin.Engine) {
	user := new(controllers.User)
	router.GET("/user", user.GetAllUsers)
	router.POST("/user/signin", user.Login)
	router.POST("/user/signup", user.RegisterUser)
	router.POST("/user/update", user.Update)
}
