package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	Username   string    `json:"username" gorm:"unique;not null"`
	Password   string    `json:"password" gorm:"not null"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Rentals    []Rentals `json:"rentals" gorm:"foreignkey:UserID"`
}

func (User) TableName() string {
	return "tbl_users"
}

// func (u *User) GetFormattedCreatedAt()string {
// 	return u.Created_at.Format("02-01-2006")
// }

// func (u *User) GetFormattedUpdatedAt()string {
// 	return u.Updated_at.Format("02-01-2006")
// }

// hash password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	return err == nil
}
