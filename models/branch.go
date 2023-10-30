package models

type Branch struct {
	Id       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Locality string `json:"locality"`
	City     string `json:"city"`
	State    string `json:"state"`
	Pincode  int64  `json:"pincode"`
}

// @Column(nullable = false)
// private LocalDateTime createdAt = LocalDateTime.now();

// @Column(nullable = false)
// private LocalDateTime modifiedAt = LocalDateTime.now();
