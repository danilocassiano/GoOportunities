package router

import (
	"github.com/danilocassiano/GoOportunities/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/danilocassiano/GoOportunities/docs"
) 

func initializeRoutes(router *gin.Engine){
	
	//initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}	

	//inicialize Swagger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
}

