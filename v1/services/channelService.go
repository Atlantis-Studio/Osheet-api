package services

import (
	"Osheet-api/database"
	"Osheet-api/v1/models"
)

type ChannelService struct{}

// Get all channels
func (s ChannelService) GetAllChannel() (channelList []*models.Channel, err error) {
	if err = database.GetDB().Find(&channelList).Error; err != nil {
		return nil, err
	}

	// TODO: database.GetDB().Find(&channelList, "company = ?", "Hololive")

	return channelList, nil
}

// Create channel
func (s ChannelService) StoreChannel(newChannel models.Channel) (channel models.Channel, err error) {
	result := database.GetDB().Create(&newChannel)

	if result.Error != nil {
		return channel, err
	}

	return newChannel, nil
}

// Get channel by twiiter account
func (s ChannelService) GetChannelByTwitterAccount(twitterAccount string) (channel *models.Channel, err error) {
	if err = database.GetDB().Where("twitter_account = ?", twitterAccount).First(&channel).Error; err != nil {
		return nil, err
	}

	return channel, nil
}
