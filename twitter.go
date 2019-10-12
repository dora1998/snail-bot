package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
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
}
