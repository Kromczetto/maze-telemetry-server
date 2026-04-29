package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	loadConfig()
	connectDB()

	r := gin.Default()
	r.Use(cors.Default())
	setupRoutes(r)
	r.Run(":" + AppConfig.Port)
}