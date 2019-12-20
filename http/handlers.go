package http

import (
	"fmt"
	"github.com/dora1998/snail-bot/twitter"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CallbackBody struct {
	Text        string `json:"text"`
	UserName    string `json:"user_name"`
	LinkToTweet string `json:"link_to_tweet"`
	CreatedAt   string `json:"created_at"`
}

func (s *Server) IFTTTCallback(c *gin.Context) {
	callbackBody := &CallbackBody{}
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
