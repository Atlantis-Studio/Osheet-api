package controllers

import (
	services "Osheet-api/v1/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChannelController struct{}

// @Summary     Get Channels
// @Description Get channels by filters
// @Tags        Channels
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} []models.Channel
// @Router 		/api/v1/channels [get]
func (c ChannelController) Index(context *gin.Context) {
	// company := context.Query("company")  // TODO get query string: company

	channelService := new(services.ChannelService)
	channels, err := channelService.GetAllChannel()
	if err != nil {
		fmt.Println("Get /api/v1/channels Failed:", err)
	}

	context.JSON(http.StatusOK, channels)
}

// @Summary     Create Channl
// @Description Store new channel
// @Tags        Channels
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} models.Channel
// @Router 		/api/v1/channels [post]
func (c ChannelController) Store(context *gin.Context) {
}

// @Summary     Get Channl By Twitter Account
// @Description Show channel info
// @Tags        Channels
// @Accept      json
// @Produce     json
// @param       twitterAccount path string true "Twitter Account"
// @Success 	200 {object} models.Channel
// @Failure		404 {string} Channel not found
// @Router 		/api/v1/channels/{twitterAccount} [get]
func (c ChannelController) Show(context *gin.Context) {
	twitterAccount := context.Param("twitterAccount") // get URL path parameter: twitterAccount

	channelService := new(services.ChannelService)
	channel, err := channelService.GetChannelByTwitterAccount(twitterAccount)
	if err != nil {
		fmt.Println("Get /api/v1/channels/{twitterAccount} Failed:", err)
	}

	context.JSON(http.StatusOK, channel)
}
