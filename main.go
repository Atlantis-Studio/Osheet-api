package main

import (
	"github.com/gin-gonic/gin"

	ChannelController "Osheet-api/Channel/Controller"

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
	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			v1Group.GET("/channels", ChannelController.Index)
			v1Group.GET("/channels/:twitterAccount", ChannelController.Show)

			// TODO POST /channels
		}
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
