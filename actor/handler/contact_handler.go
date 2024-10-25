package handler

import "github.com/gin-gonic/gin"

type ContactInput struct {
	ActorId      int    `json:"actor_id" binding:"required"`
	ContactType  string `json:"contact_type" binding:"required"`
	ContactValue string `json:"contact_value" binding:"required"`
}

type ContactOutput struct{}

func GetContactByID(c *gin.Context) {

}

func CreateContact(c *gin.Context) {

}

func UpdateContact(c *gin.Context) {}

func DeactivateContact(c *gin.Context) {}
