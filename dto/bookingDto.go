package dto

import "time"

type BookingAddDto struct {
	TripCost      int64     `json:"tripCost"`
	BookingDate   time.Time `json:"bookingDate" `
	BookingStatus string    `json:"bookingStatus"`
	FromDate      time.Time `json:"fromDate"`
	ToDate        time.Time `json:"toDate"`
	BranchID      int64     `json:"branchID"`
	UserID        int64     `json:"userID"`
	CarID         int64     `json:"carID"`

	TransId     string `json:"transId"`
	Amount      int64  `json:"amount"`
	PaymentMode string `json:"paymentMode"`
}
