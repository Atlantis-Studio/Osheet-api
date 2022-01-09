package models

import (
	"gorm.io/gorm"
)

// TODO: FanName, Hashtag(StreamTag, FanArt)

type Channel struct {
	gorm.Model
	Name           string  `json:"name" gorm:"not null"`
	TwitterAccount string  `json:"twitterAccount" gorm:"index;unique;not null"`
	Avatar         string  `json:"avatar" gorm:"default:null"`
	Company        string  `json:"company" gorm:"default:null"`
	Unit           string  `json:"unit" gorm:"default:null"`
	ChannelUrl     string  `json:"channelUrl" gorm:"not null"`
	Birthday       string  `json:"birthday" gorm:"default:null"`
	Height         float64 `json:"height" gorm:"default:null"`
	DebutDate      string  `json:"debutDate" gorm:"default:null"`
}
