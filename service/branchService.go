package service

import (
	"github.com/harshalhonmote/crs/database"
	"github.com/harshalhonmote/crs/models"
	"gorm.io/gorm"
)

type Branch struct {
	database.Database
}

func (s *Branch) Create(branch *models.Branch) (*gorm.DB, *models.Branch) {
	result := s.GetDB().Create(&branch)
	return result, branch
}
