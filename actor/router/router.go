package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jpcairesf/cinema/actor/handler"
)

func Initialize() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// Include routeGroup for actor and test
	r.GET("/actor", handler.GetActors)
	r.GET("/actor/{id}", handler.GetActorByID)
	r.POST("/actor", handler.CreateActor)
	r.PUT("/actor", handler.UpdateActor)
	r.DELETE("/actor/{id}", handler.DeleteActor)
	// Include routeGroup for actor/{id}/contact and test
	r.GET("/actor/{id}/contact", func(c *gin.Context) {})
	r.POST("/actor/{id}/contact", func(c *gin.Context) {})
	r.PUT("/actor/{id}/contact", func(c *gin.Context) {})
	r.PATCH("/actor/{id}/contact", func(c *gin.Context) {})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
