package entity

import (
	"gorm.io/gorm"
)

type Actor struct {
	gorm.Model
	FirstName string
	LastName  string
	//	BirthDate   time.Time
	Nationality string
	Gender      string
	IsRetired   bool
	Contacts    []Contact
}

type ActorOutput struct {
}
