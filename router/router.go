package router

import "github.com/gin-gonic/gin" 

func Inicialize(){

	//Initializer router.
	router := gin.Default()

	initializeRoutes(router)
	
	
	//Run the server.
	router.Run(":8080")
}
