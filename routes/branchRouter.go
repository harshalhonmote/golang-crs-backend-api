package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/controllers"
)

func BranchRoutes(router *gin.Engine) {
	branch := new(controllers.Branch)
	router.POST("/branch/add", branch.AddBranch)
	// router.POST("/user/signin", user.Login)
	// router.POST("/user/signup", user.RegisterUser)
	// router.POST("/user/update", user.Update)
}
