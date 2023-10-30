package models

import "time"

type Transaction struct {
	BookingId int64    `json:"id" gorm:"unique;not null"`
	Booking   *Booking `gorm:"foreignKey:BookingId"`

	TransId     string    `json:"transId"`
	Amount      int64     `json:"amount"`
	PaymentMode string    `json:"paymentMode"`
	CreatedAt   time.Time `json:"createdAt"`
}
