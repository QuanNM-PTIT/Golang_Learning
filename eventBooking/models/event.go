package models

import (
	"eventBooking/database"
)

type Event struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	Location    string `binding:"required" json:"location"`
	Datetime    string `json:"datetime"`
	UserID      int    `gorm:"column:user_id" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"user"`
}

func (e Event) Save() error {
	result := database.DB.Create(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetEvents() ([]Event, error) {
	var events []Event
	result := database.DB.Find(&events)
	if result.Error != nil {
		return events, result.Error
	}
	return events, nil
}

func GetEventById(id string) (Event, error) {
	var event Event
	result := database.DB.First(&event, id)
	if result.Error != nil {
		return event, result.Error
	}
	return event, nil
}

func (e Event) Update() error {
	result := database.DB.Updates(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (e Event) Delete() error {
	result := database.DB.Delete(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
