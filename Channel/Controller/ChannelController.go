package ChannelController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ChannelEntity "Osheet-api/Channel/Entity"
)

func Index(context *gin.Context) {
	// company := context.Query("company")  // get query string: company

	ina := ChannelEntity.Channel{
		TwitterAccount: "ninomaeinanis",
		Name:           "Ninomae Ina'nis",
		Avatar:         "https://pbs.twimg.com/profile_images/1473599637751693314/k_ZTGRX4_400x400.jpg",
		Company:        "Hololive",
		Unit:           "hololive English - Myth",
		ChannelUrl:     "https://www.youtube.com/channel/UCMwGHR0BTZuLsmjY_NT5Pwg",
		Height:         "157",
		Birthday:       "05/20",
		DebutDate:      "2020/09/13",
		Info: ChannelEntity.Info{
			FanName: "takodachi tentacult",
			Hashtag: ChannelEntity.Hashtag{
				StreamTag: "#TAKOTIME #タコタイム",
				FanArt:    "#inART #いなート",
			},
		},
	}

	ame := ChannelEntity.Channel{
		TwitterAccount: "watsonameliaEN",
		Name:           "Watson Amelia",
		Avatar:         "https://pbs.twimg.com/profile_images/1318958836120649728/7JHxy2UO_400x400.jpg",
		Company:        "Hololive",
		Unit:           "hololive English - Myth",
		ChannelUrl:     "https://www.youtube.com/channel/UCyl1z3jo3XHR1riLFKG5UAg",
		Height:         "150",
		Birthday:       "01/06",
		DebutDate:      "2020/09/13",
		Info: ChannelEntity.Info{
			FanName: "teamates",
			Hashtag: ChannelEntity.Hashtag{
				StreamTag: "#amelive",
				FanArt:    "#ameliaRT",
			},
		},
	}

	var channels []ChannelEntity.Channel
	channels = append(channels, ina, ame)

	context.JSON(http.StatusOK, channels)
}

func Show(context *gin.Context) {
	twitterAccount := context.Param("twitterAccount") // get URL path parameter: twitterAccount

	channel := ChannelEntity.Channel{
		TwitterAccount: twitterAccount,
		Name:           "Garw Gura",
		Avatar:         "https://pbs.twimg.com/profile_images/1309957523089354760/uRrxAmOB_400x400.jpg",
		Company:        "Hololive",
		Unit:           "hololive English - Myth",
		ChannelUrl:     "https://www.youtube.com/channel/UCoSrY_IQQVpmIRZ9Xf-y93g",
		Height:         "141",
		Birthday:       "06/20",
		DebutDate:      "2020/09/13",
		Info: ChannelEntity.Info{
			FanName: "chumbuds",
			Hashtag: ChannelEntity.Hashtag{
				StreamTag: "#gawrgura",
				FanArt:    "#gawrt",
			},
		},
	}

	context.JSON(http.StatusOK, channel)
}
