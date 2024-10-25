package handler

import "github.com/gin-gonic/gin"

type ActorInput struct {
	FirstName   string         `json:"first_name" binding:"required" maxLength:"32"`
	LastName    string         `json:"last_name" binding:"required" maxLength:"64"`
	Nationality string         `json:"nationality" binding:"required" maxLength:"32"`
	Gender      string         `json:"gender" binding:"required"`
	IsRetired   bool           `json:"is_retired" binding:"required"`
	Contacts    []ContactInput `json:"contacts"`
}

type ActorOutput struct {
}

func GetActors(c *gin.Context) {

}

func GetActorByID(c *gin.Context) {

}

func CreateActor(c *gin.Context) {

}

func UpdateActor(c *gin.Context) {

}

func DeleteActor(c *gin.Context) {}
