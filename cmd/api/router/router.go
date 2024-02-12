package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initializing Router with gin default config
	r := gin.Default()

	//Defining route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Running API server (8080 default port)
	r.Run(":8080")
}
