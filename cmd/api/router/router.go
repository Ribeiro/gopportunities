package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initializing Router with gin default config
	r := gin.Default()

	//Initialize Routes
	initializeRoutes(r)

	//Running API server (8080 default port)
	r.Run(":8080")
}
