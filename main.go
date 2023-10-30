package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/database"
	"github.com/harshalhonmote/crs/routes"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	r := gin.Default()
	database.Migrate()

	routes.UserRoutes(r)
	routes.BranchRoutes(r)
	routes.CarRoutes(r)
	routes.BookingRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3001", r))
	r.Run()
}
