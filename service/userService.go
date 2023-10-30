package service

import (
	"github.com/harshalhonmote/crs/database"
	"github.com/harshalhonmote/crs/models"
	"gorm.io/gorm"
)

type User struct {
	database.Database
}

func (s *User) Create(user *models.User) (*gorm.DB, *models.User) {
	result := s.GetDB().Create(&user)
	return result, user
}
func (s *User) GetAllUsers() (*gorm.DB, []models.User) {
	var user []models.User
	result := s.GetDB().Find(&user)
	return result, user
}

func (s *User) Login(email string, password string) (*gorm.DB, models.User) {
	var user models.User
	result := s.GetDB().Where("email = ? and password = ?", email, password).First(&user)
	return result, user
}
func (s *User) Update(user *models.User) (*gorm.DB, *models.User) {
	result := s.GetDB().Save(&user)

	//result := s.GetDB().Model(&user).Select("FirstName").Updates(models.User{FirstName: user.FirstName})
	return result, user
}
