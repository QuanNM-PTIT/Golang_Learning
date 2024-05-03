package models

import (
	"errors"
	"eventBooking/database"
	"eventBooking/utils"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Email    string `binding:"required" json:"email" gorm:"unique index"`
	Password string `binding:"required" json:"password"`
}

func (u User) Save() error {
	result := database.DB.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) ValidateCredentials() error {
	email, password := u.Email, u.Password
	result := database.DB.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return errors.New("Invalid credentials!")
	}
	if !utils.ComparePassword(u.Password, password) {
		return errors.New("Invalid credentials!")
	}
	return nil
}
