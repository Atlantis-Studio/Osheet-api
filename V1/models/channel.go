package models

import (
	"gorm.io/gorm"
)

// TODO: FanName, Hashtag(StreamTag, FanArt)

type Channel struct {
	gorm.Model
	Name           string  `json:"name"`
	TwitterAccount string  `json:"twitterAccount"`
	Avatar         string  `json:"avatar"`
	Company        string  `json:"company"`
	Unit           string  `json:"unit"`
	ChannelUrl     string  `json:"channelUrl"`
	Birthday       string  `json:"birthday"`
	Height         float64 `json:"height"`
	DebutDate      string  `json:"debutDate"`
}
