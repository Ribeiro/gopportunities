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
	//Initialize Handler
	handler.InitializeHandler()

	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		v1.GET(openingsPath, handler.GetOpeningsHandler)
		v1.GET(openingsPath+"/:id", handler.GetOpeningHandler)
		v1.POST(openingsPath, handler.CreateOpeningHandler)
		v1.DELETE(openingsPath, handler.DeleteOpeningHandler)
		v1.PUT(openingsPath, handler.UpdateOpeningHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
