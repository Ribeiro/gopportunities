package router

import (
	"github.com/Ribeiro/gopportunities/cmd/api/handler"
	docs "github.com/arthur404dev/gopportunities/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const basePath = "/api/v1"
const openingsPath = "/openings"

func initializeRoutes(router *gin.Engine) {

	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		v1.GET(openingsPath, handler.GetAllOpeningsHandler)
		v1.GET(openingsPath+"/:id", handler.GetOpeningsHandler)
		v1.POST(openingsPath, handler.CreateOpeningsHandler)
		v1.DELETE(openingsPath, handler.DeleteOpeningsHandler)
		v1.PUT(openingsPath, handler.UpdateOpeningsHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
