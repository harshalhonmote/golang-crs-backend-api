package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harshalhonmote/crs/helper"
	"github.com/harshalhonmote/crs/models"
	"github.com/harshalhonmote/crs/service"
)

type Branch struct {
	helper.Controller
}

func (car *Branch) AddBranch(c *gin.Context) {
	var params *models.Branch
	c.BindJSON(&params)
	result, data := new(service.Branch).Create(params)
	car.Json(c, result, data)
}
