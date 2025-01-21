package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

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
