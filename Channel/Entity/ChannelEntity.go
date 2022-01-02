package ChannelEntity

type Hashtag struct {
	StreamTag string `json:"streamTag"`
	FanArt    string `json:"fanArt"`
}

type Info struct {
	FanName string  `json:"fanName"`
	Hashtag Hashtag `json:"hashtags"`
}

type Channel struct {
	TwitterAccount string  `json:"twitterAccount"`
	Name           string  `json:"name"`
	Avatar         string  `json:"avatar"`
	Company        string  `json:"company"`
	Unit           string  `json:"unit"`
	ChannelUrl     string  `json:"channelUrl"`
	Birthday       string  `json:"birthday"`
	Height         float64 `json:"height"`
	DebutDate      string  `json:"debutDate"`
	Info           Info    `json:"info"`
}
