package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
)

func tweet(msg string) {
	var env Env
	err := envconfig.Process("bot", &env)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", env)

	config := oauth1.NewConfig(env.ConsumerKey, env.ConsumerSecret)
	token := oauth1.NewToken(env.Token, env.TokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, _, err := client.Statuses.Update(msg, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n", tweet)
}
