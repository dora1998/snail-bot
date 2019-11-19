package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dora1998/snail-bot/utils"
	"github.com/kelseyhightower/envconfig"
	"regexp"
	"strconv"
)

type TwitterClient struct {
	client *twitter.Client
}

func NewTwitterClient() *TwitterClient {
	var env utils.Env
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

func (c *TwitterClient) Tweet(msg string) *twitter.Tweet {
	// Send a Tweet
	tweet, _, err := c.client.Statuses.Update(msg, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n", tweet)
	return tweet
}

func (c *TwitterClient) Reply(msg string, tweetId int64) *twitter.Tweet {
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

func (c *TwitterClient) CreateFavorite(tweetId int64) error {
	_, _, err := c.client.Favorites.Create(&twitter.FavoriteCreateParams{ID: tweetId})
	return err
}

func (c *TwitterClient) IsFollwing(screenName string) bool {
	user, _, err := c.client.Users.Show(&twitter.UserShowParams{
		ScreenName: screenName,
	})

	if user == nil || err != nil {
		return false
	}

	return user.Following
}

func ExtractStatusIdFromUrl(url string) (int64, error) {
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

func ExtractBody(text string) (string, error) {
	regexpObj := regexp.MustCompile(".*@assignment_bot (.+)$")
	res := regexpObj.FindStringSubmatch(text)

	if res == nil {
		return "", fmt.Errorf("failed extractBody (no match)")
	}
	return res[1], nil
}
