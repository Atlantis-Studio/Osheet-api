package main

import (
	"github.com/gin-gonic/gin"

	ChannelController "Osheet-api/Channel/Controller"
)

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

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
