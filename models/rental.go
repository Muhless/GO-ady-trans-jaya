package models

import "time"

type Rentals struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	CarID      uint      `json:"car_id" gorm:"not null"`
	Car        Cars      `json:"cars" gorm:"foreignkey:CarID;contsraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	User       User      `json:"tbl_users" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status     string    `json:"status" gorm:"size:50;not null"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Rentals) TableName() string {
	return "tbl_rentals"
}
