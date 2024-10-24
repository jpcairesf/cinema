package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// Include routeGroup for actor and test
	r.GET("/actor", func(c *gin.Context) {})
	r.GET("/actor/{id}", func(c *gin.Context) {})
	r.POST("/actor", func(c *gin.Context) {})
	r.PUT("/actor", func(c *gin.Context) {})
	r.DELETE("/actor/{id}", func(c *gin.Context) {})
	// Include routeGroup for actor/{id}/contact and test
	r.GET("/actor/{id}/contact", func(c *gin.Context) {})
	r.POST("/actor/{id}/contact", func(c *gin.Context) {})
	r.PUT("/actor/{id}/contact", func(c *gin.Context) {})
	r.DELETE("/actor/{id}/contact/{id}", func(c *gin.Context) {})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
