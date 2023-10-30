package models

import "time"

type Booking struct {
	Id            int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingStatus string `json:"bookingStatus"`

	TripCost int64 `json:"tripCost"`
	Penalty  int64 `json:"penalty"  gorm:"default:0"`

	BookingDate time.Time `json:"bookingDate"`
	FromDate    time.Time `json:"fromDate"`
	ToDate      time.Time `json:"toDate"`

	BranchID int64   `json:"branchID" gorm:"not null"`
	Branch   *Branch `json:"branch"`

	UserID int64 `json:"userID" gorm:"not null"`
	User   *User `json:"User"`

	CarID int64 `json:"carID" gorm:"not null"`
	Car   *Car  `json:"car"`

	// CompanyRefer int
	// Company      Company `gorm:"foreignKey:CompanyRefer"`
}
