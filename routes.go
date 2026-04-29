package main

import "github.com/gin-gonic/gin"

func setupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/run", saveRun)

	r.GET("/runs", getAllRuns)
}