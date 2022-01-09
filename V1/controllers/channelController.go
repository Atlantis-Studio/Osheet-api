package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	service "Osheet-api/v1/services"
)

type ChannelController struct{}

// @Summary     Get Channels
// @Description Get channels by filters
// @Tags        Channels
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} []ChannelEntity.Channel
// @Router 		/api/v1/channels [get]
func (c ChannelController) Index(context *gin.Context) {
	// company := context.Query("company")  // TODO get query string: company

	channels, err := service.GetAllChannel()
	if err != nil {
		fmt.Println("Get /api/v1/channels Failed:", err)
	}

	context.JSON(http.StatusOK, channels)
}

// @Summary Get Channel By Twitter Account
// @Description Show channel info
// @Tags        Channels
// @Accept      json
// @Produce     json
// @param       twitterAccount path string true "Twitter Account"
// @Success 	200 {object} ChannelEntity.Channel
// @Failure		404 {string} Channel not found
// @Router 		/api/v1/channels/{twitterAccount} [get]
func (c ChannelController) Show(context *gin.Context) {
	twitterAccount := context.Param("twitterAccount") // get URL path parameter: twitterAccount

	channel, err := service.GetChannelByTwitterAccount(twitterAccount)
	if err != nil {
		fmt.Println("Get /api/v1/channels/{twitterAccount} Failed:", err)
	}

	context.JSON(http.StatusOK, channel)
}

// TODO: Create channel
