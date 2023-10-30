# golang-crs-backend-api
- same backend created in java spring boot you can checkout [here](https://github.com/Online-Car-Rental-System-CDAC-Proj/BackEnd)
- here i used GORM which is Object Relational Mapping (ORM) library for Golang, similer to Hibernate in Java.
- Gin is a high-performance HTTP web framework written in Golang (Go). (Fiber web framework also populer one)
  To Run
  ```
    $ go mod download
    $ go run main.go
  ```
#### Model Example in Golang
```
type User struct {
	Id            int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email" gorm:"unique;not null"`
	Password      string `json:"password"`
	KycStatus     bool   `json:"kycstatus" gorm:"default:false"`
	AccountStatus bool   `json:"accountstatus" gorm:"default:false"`
}

type Branch struct {
	Id       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Locality string `json:"locality"`
	City     string `json:"city"`
	State    string `json:"state"`
	Pincode  int64  `json:"pincode"`
}

type Car struct {
	Id           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	Availability bool      `json:"availability" gorm:"default:true"`
	ImageUrl     string    `json:"imageUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	BranchID     int64     `json:"branchID" gorm:"not null"`
	Branch       *Branch   `json:"branch"`
}

type Booking struct {
	Id            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingStatus string    `json:"bookingStatus"`
	TripCost      int64     `json:"tripCost"`
	Penalty       int64     `json:"penalty"  gorm:"default:0"`
	BookingDate   time.Time `json:"bookingDate"`
	FromDate      time.Time `json:"fromDate"`
	ToDate        time.Time `json:"toDate"`
	BranchID      int64     `json:"branchID" gorm:"not null"`
	Branch        *Branch   `json:"branch"`
	UserID        int64     `json:"userID" gorm:"not null"`
	User          *User     `json:"User"`
	CarID         int64     `json:"carID" gorm:"not null"`
	Car           *Car      `json:"car"`
	// CompanyRefer int
	// Company      Company `gorm:"foreignKey:CompanyRefer"`
}
```
  
