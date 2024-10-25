package handler

import "github.com/gin-gonic/gin"

type CreateActorInput struct {
	FirstName   string               `json:"first_name" maxLength:"32"`
	LastName    string               `json:"last_name" maxLength:"64"`
	Nationality string               `json:"nationality" maxLength:"32"`
	Gender      string               `json:"gender"`
	IsRetired   bool                 `json:"is_retired"`
	Contacts    []CreateContactInput `json:"contacts"`
}

type CreateContactInput struct {
	ContactType  string `json:"contact_type" maxLength:"32"`
	ContactValue string `json:"contact_value" maxLength:"64"`
}

type UpdateActorInput struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name" maxLength:"32"`
	LastName  string `json:"last_name" maxLength:"64"`
	Gender    string `json:"gender"`
	IsRetired bool   `json:"is_retired"`
}

type ActorOutput struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name" maxLength:"32"`
	LastName    string `json:"last_name" maxLength:"64"`
	Nationality string `json:"nationality" maxLength:"32"`
	Gender      string `json:"gender"`
	IsRetired   bool   `json:"is_retired"`
}

func GetActors(c *gin.Context) {

}

func GetActorByID(c *gin.Context) {

}

func CreateActor(c *gin.Context) {
	request := CreateActorInput{}
	c.BindJSON(&request)
}

func UpdateActor(c *gin.Context) {
	request := UpdateActorInput{}
	c.BindJSON(&request)
}

func DeleteActor(c *gin.Context) {}
