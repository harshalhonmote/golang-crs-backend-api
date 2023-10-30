package models

import "time"

type Car struct {
	Id           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	Availability bool      `json:"availability" gorm:"default:true"`
	ImageUrl     string    `json:"imageUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	BranchID     int64     `json:"branchID" gorm:"not null"`
	Branch       *Branch   `json:"branch"`

	// CompanyRefer int
	// Company      Company `gorm:"foreignKey:CompanyRefer"`
}

// @ManyToOne(fetch = FetchType.LAZY)
// @JoinColumn(name = "branch_id", nullable = false)
// private Branch branch;
