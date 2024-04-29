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
	UserID      int    `gorm:"column:userID" json:"userID"`
}

func (e Event) Save() {
	database.DB.Create(&e)
}

func GetEvents() []Event {
	var events []Event
	database.DB.Find(&events)
	return events
}
