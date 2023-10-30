package controllers

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/helper"
	"github.com/harshalhonmote/crs/models"
	"github.com/harshalhonmote/crs/service"
)

var user models.User

type User struct {
	helper.Controller
}

func (u *User) Login(c *gin.Context) {
	bodyAsByteArray, _ := io.ReadAll(c.Request.Body)
	jsonMap := make(map[string]interface{})
	json.Unmarshal(bodyAsByteArray, &jsonMap)

	email := jsonMap["email"].(string)
	password := jsonMap["password"].(string)
	result, data := new(service.User).Login(email, password)
	u.Json(c, result, data)
}

func (u *User) GetAllUsers(c *gin.Context) {
	result, data := new(service.User).GetAllUsers()
	u.Json(c, result, data)
}
func (u *User) RegisterUser(c *gin.Context) {
	var params *models.User
	c.BindJSON(&params)
	result, data := new(service.User).Create(params)
	u.Json(c, result, data)
	// if result.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"data": nil,
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": data,
	// })
}

func (u *User) Update(c *gin.Context) {
	var params *models.User
	c.BindJSON(&params)
	result, data := new(service.User).Update(params)
	u.Json(c, result, data)
}
