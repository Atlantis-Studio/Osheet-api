package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	type Hashtag struct {
		StreamTag string `json:"streamTag"`
		FanArt    string `json:"fanArt"`
	}

	type Info struct {
		FanName string  `json:"fanName"`
		Hashtag Hashtag `json:"hashtags"`
	}

	type Channel struct {
		TwitterAccount string `json:"twitterAccount"`
		Name           string `json:"name"`
		Avatar         string `json:"avatar"`
		Company        string `json:"company"`
		Unit           string `json:"unit"`
		ChannelUrl     string `json:"channelUrl"`
		Birthday       string `json:"birthday"`
		Height         string `json:"height"`
		DebutDate      string `json:"debutDate"`
		Info           Info   `json:"info"`
	}

	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			v1Group.GET("/query", func(c *gin.Context) {
				company := c.Query("company")
				c.JSON(http.StatusOK, company)
			})

			v1Group.GET("/channels", func(c *gin.Context) {
				ina := Channel{
					TwitterAccount: "ninomaeinanis",
					Name:           "Ninomae Ina'nis",
					Avatar:         "https://pbs.twimg.com/profile_images/1473599637751693314/k_ZTGRX4_400x400.jpg",
					Company:        "Hololive",
					Unit:           "hololive English - Myth",
					ChannelUrl:     "https://www.youtube.com/channel/UCMwGHR0BTZuLsmjY_NT5Pwg",
					Height:         "157",
					Birthday:       "05/20",
					DebutDate:      "2020/09/13",
					Info: Info{
						FanName: "takodachi tentacult",
						Hashtag: Hashtag{
							StreamTag: "#TAKOTIME #タコタイム",
							FanArt:    "#inART #いなート",
						},
					},
				}

				ame := Channel{
					TwitterAccount: "watsonameliaEN",
					Name:           "Watson Amelia",
					Avatar:         "https://pbs.twimg.com/profile_images/1318958836120649728/7JHxy2UO_400x400.jpg",
					Company:        "Hololive",
					Unit:           "hololive English - Myth",
					ChannelUrl:     "https://www.youtube.com/channel/UCyl1z3jo3XHR1riLFKG5UAg",
					Height:         "150",
					Birthday:       "01/06",
					DebutDate:      "2020/09/13",
					Info: Info{
						FanName: "teamates",
						Hashtag: Hashtag{
							StreamTag: "#amelive",
							FanArt:    "#ameliaRT",
						},
					},
				}

				var channels []Channel
				channels = append(channels, ina, ame)

				c.JSON(http.StatusOK, channels)
			})

			v1Group.GET("/channels/:twitterAccount", func(c *gin.Context) {
				twitterAccount := c.Param("twitterAccount") // 取得 URL 路徑中參數 twitterAccount

				channel := Channel{
					TwitterAccount: twitterAccount,
					Name:           "Garw Gura",
					Avatar:         "https://pbs.twimg.com/profile_images/1309957523089354760/uRrxAmOB_400x400.jpg",
					Company:        "Hololive",
					Unit:           "hololive English - Myth",
					ChannelUrl:     "https://www.youtube.com/channel/UCoSrY_IQQVpmIRZ9Xf-y93g",
					Height:         "141",
					Birthday:       "06/20",
					DebutDate:      "2020/09/13",
					Info: Info{
						FanName: "chumbuds",
						Hashtag: Hashtag{
							StreamTag: "#gawrgura",
							FanArt:    "#gawrt",
						},
					},
				}

				c.JSON(http.StatusOK, channel)
			})

			// TODO POST /channels
		}
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
