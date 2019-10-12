package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
	"regexp"
	"strconv"
)

type TwitterClient struct {
	client *twitter.Client
}

func NewTwitterClient() *TwitterClient {
	var env Env
	err := envconfig.Process("bot", &env)
	if err != nil {
		fmt.Println(err.Error())
	}

	config := oauth1.NewConfig(env.ConsumerKey, env.ConsumerSecret)
	token := oauth1.NewToken(env.Token, env.TokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return &TwitterClient{client: client}
}

func (c *TwitterClient) tweet(msg string) *twitter.Tweet {
	// Send a Tweet
	tweet, _, err := c.client.Statuses.Update(msg, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n", tweet)
	return tweet
}

func (c *TwitterClient) reply(msg string, tweetId int64) *twitter.Tweet {
	// Send a Tweet
	tweet, _, err := c.client.Statuses.Update(msg, &twitter.StatusUpdateParams{
		InReplyToStatusID:         tweetId,
		AutoPopulateReplyMetadata: Bool(true),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n", tweet)
	return tweet
}

func extractStatusIdFromUrl(url string) (int64, error) {
	regexpObj := regexp.MustCompile("^http://twitter.com/.+/status/(.+)$")
	res := regexpObj.FindStringSubmatch(url)

	if res == nil {
		return 0, fmt.Errorf("failed extractStatusIdFromUrl (no match)")
	}

	id, err := strconv.ParseInt(res[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
