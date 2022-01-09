package main

import (
	"github.com/gin-gonic/gin"

	"Osheet-api/database"
	controllers "Osheet-api/v1/controllers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Osheet-api/docs"
)

// @title Oshi-api
// @version 1.0
// @description Oshi-api swagger

// @license.name MIT
// @license.url https://www.mit.edu/~amini/LICENSE.md

// schemes http
func main() {
	defer database.Close()

	router := gin.Default()

	channelController := new(controllers.ChannelController)

	apiGroup := router.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			v1Group.GET("/channels", channelController.Index)
			v1Group.GET("/channels/:twitterAccount", channelController.Show)

			// TODO POST /channels
		}
	}

	url := ginSwagger.URL("http://0.0.0.0:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
