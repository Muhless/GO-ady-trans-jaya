package models

import "time"

type Cars struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	Category     string    `json:"category"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	Availability bool      `json:"availability"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Rentals      []Rentals `json:"rentals" gorm:"foreignkey:CarID"`
}

func (Cars) TableName() string {
	return "tbl_cars"
}
