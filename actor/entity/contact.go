package entity

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ActorId      int
	ContactType  string
	ContactValue string
	IsActive     bool
}

type ContactOutput struct {
}
