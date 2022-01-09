package service

import (
	"Osheet-api/database"
	"Osheet-api/v1/models"
)

// Get all channels
func GetAllChannel() (channelList []*models.Channel, err error) {
	if err = database.GetDB().Find(&channelList).Error; err != nil {
		return nil, err
	}

	return channelList, nil
}

// Get channel by twiiter account
func GetChannelByTwitterAccount(twitterAccount string) (channel *models.Channel, err error) {
	if err = database.GetDB().Where("twitter_account = ?", twitterAccount).First(&channel).Error; err != nil {
		return nil, err
	}

	return channel, nil
}
