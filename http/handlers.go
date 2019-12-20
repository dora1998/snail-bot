package http

import (
	"fmt"
	_twitter "github.com/dghubble/go-twitter/twitter"
	"github.com/dora1998/snail-bot/twitter"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostIFTTTWebHookRequest struct {
	Text        string `json:"text"`
	UserName    string `json:"user_name"`
	LinkToTweet string `json:"link_to_tweet"`
	CreatedAt   string `json:"created_at"`
}

func (s *Server) PostIFTTTWebHook(c *gin.Context) {
	callbackBody := &PostIFTTTWebHookRequest{}
	err := c.BindJSON(&callbackBody)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("%#v\n", callbackBody)
	statusId, err := twitter.ExtractStatusIdFromUrl(callbackBody.LinkToTweet)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	text, err := twitter.ExtractBody(callbackBody.Text)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.commandHandler.Resolve(text, callbackBody.UserName, statusId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (s *Server) GetCRCToken(c *gin.Context) {
	crcToken := c.Query("crc_token")
	resToken := s.twitterClient.CreateCRCToken(crcToken)
	c.JSON(http.StatusOK, gin.H{"response_token": resToken})
}

type PostTwitterWebHookRequest struct {
	TweetCreateEvents []_twitter.Tweet `json:"tweet_create_events"`
}

func (s *Server) PostWebHook(c *gin.Context) {
	body := &PostTwitterWebHookRequest{}
	err := c.BindJSON(body)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("%#v\n", body)

	if len(body.TweetCreateEvents) == 0 {
		c.String(http.StatusOK, "")
		return
	}

	for _, t := range body.TweetCreateEvents {
		// 自分宛の@ツイート以外は無視(引用RT, Botが返信した自身のツイートなどは除外)
		if t.InReplyToScreenName != "assignment_bot" {
			continue
		}

		statusId := t.ID

		text, err := twitter.ExtractBody(t.Text)
		if err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusOK, "")
			return
		}

		err = s.commandHandler.Resolve(text, t.User.Name, statusId)
		if err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusOK, "")
			return
		}
	}

	c.String(http.StatusOK, "")
}
