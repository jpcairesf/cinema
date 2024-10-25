package handler

import "github.com/gin-gonic/gin"

type AddContactInput struct {
	ActorId      int    `json:"actor_id"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
}

type ContactOutput struct {
	Id           int    `json:"id"`
	ActorId      int    `json:"actor_id"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
}

func GetContactsByActorID(c *gin.Context) {

}

func AddContactByActorID(c *gin.Context) {
	request := AddContactInput{}
	c.BindJSON(&request)
}

func DeactivateContact(c *gin.Context) {}
