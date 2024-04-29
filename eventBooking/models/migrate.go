package models

import "eventBooking/database"

func Migrate() {
	err := database.DB.AutoMigrate(&Event{})

	if err != nil {
		panic("Failed to migrate table!")
	}
}
