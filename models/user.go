package models

//var db *gorm.DB

type User struct {
	Id            int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email" gorm:"unique;not null"`
	Password      string `json:"password"`
	KycStatus     bool   `json:"kycstatus" gorm:"default:false"`
	AccountStatus bool   `json:"accountstatus" gorm:"default:false"`
}

// func init() {
// 	db = config.GetDB()
// 	db.AutoMigrate(&User{})
// }

// func (b *User) CreateBook() {
// 	//db.NewRecord(b)
// 	db.Create(&b)
// }

// func GetBook() []User {
// 	var books []User
// 	db.Find(&books)
// 	return books
// }
