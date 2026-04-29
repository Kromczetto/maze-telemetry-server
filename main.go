package main 

import github.com/gin-gonic/gin

func main() {
	loadConfig()
	connectDb()

	r := gin.Default()

	setupRoutes(r)
	r.Run(":" + AppConfig.Port)
}