package controllers

import (
	models "Osheet-api/v1/models"
	services "Osheet-api/v1/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChannelController struct{}

// @Summary     Get Channels
// @Description Get channels by filters
// @Tags        Channels
// @Accept 		json
// @Produce 	json
//
// @Success 	200 {object} []models.Channel
//
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
//
// @param       name formData string true "Display Name"
// @param       twitterAccount formData string true "Twitter Account"
// @param       avatar formData string false "Avatar Photo Url"
// @param       company formData string false "Company"
// @param       unit formData string false "Unit"
// @param       channelUrl formData string true "Channel Url"
// @param       birthday formData string false "Birthday"
// @param       height formData float64 false "Height"
// @param       debutDate formData string false "Debut Date"
//
// @Success 	200 {object} models.Channel
//
// @Router 		/api/v1/channels [post]
func (c ChannelController) Store(context *gin.Context) {
	var channel models.Channel
	channel.Name = context.PostForm("name")
	channel.TwitterAccount = context.PostForm("twitterAccount")
	channel.Avatar = context.PostForm("avatar")
	channel.Company = context.PostForm("company")
	channel.Unit = context.PostForm("unit")
	channel.ChannelUrl = context.PostForm("channelUrl")
	channel.Birthday = context.PostForm("birthday")

	if height, err := strconv.ParseFloat(context.PostForm("height"), 64); err == nil {
		channel.Height = float64(height)
	}

	channel.DebutDate = context.PostForm("debutDate")

	channelService := new(services.ChannelService)
	// TODO: check channel exist

	newChannel, err := channelService.StoreChannel(channel)
	if err != nil {
		fmt.Println("Post /api/v1/channels Failed:", err)
	}

	context.JSON(http.StatusOK, newChannel)
}

// @Summary     Get Channel By Twitter Account
// @Description Show channel info
// @Tags        Channels
// @Accept      json
// @Produce     json
//
// @param       twitterAccount path string true "Twitter Account"
//
// @Success 	200 {object} models.Channel
// @Failure		404 {string} Channel not found
//
// @Router 		/api/v1/channels/{twitterAccount} [get]
func (c ChannelController) Show(context *gin.Context) {
	twitterAccount := context.Param("twitterAccount") // get URL path parameter: twitterAccount

	channelService := new(services.ChannelService)
	channel, err := channelService.GetChannelByTwitterAccount(twitterAccount)
	if err != nil {
		fmt.Println("Get /api/v1/channels/{twitterAccount} Failed:", err)
	}

	if channel == nil {
		context.JSON(http.StatusNotFound, fmt.Sprintf("Channel %s Not Found", twitterAccount))
		return
	}

	context.JSON(http.StatusOK, channel)
}
