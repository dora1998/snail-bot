package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dora1998/snail-bot/utils"
	"github.com/kelseyhightower/envconfig"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const TWEET_MAX_LENGTH = 140

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

func (c *TwitterClient) Tweet(msg string) (*twitter.Tweet, error) {
	// Send a Tweet
	tweet, _, err := c.client.Statuses.Update(msg, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", tweet)
	return tweet, nil
}

func (c *TwitterClient) TweetLongText(text string, headText string) ([]*twitter.Tweet, error) {
	// TODO: テンプレート文字列を解析後の文字数で計算する
	maxLength := TWEET_MAX_LENGTH - utf8.RuneCountInString(headText)
	texts := SplitLongText(text, maxLength)

	pages := len(texts)
	tweets := make([]*twitter.Tweet, pages)
	var prevStatusId int64 = 0
	for i, t := range texts {
		text := headText + "\n"
		text = strings.ReplaceAll(text, "{paged}", strconv.Itoa(i+1))
		text = strings.ReplaceAll(text, "{pages}", strconv.Itoa(pages))
		text += t
		if i == 0 {
			tweet, err := c.Tweet(text)
			if err != nil {
				return tweets, err
			}
			tweets[i] = tweet
			prevStatusId = tweet.ID
		} else {
			tweet, err := c.Reply(text, prevStatusId)
			if err != nil {
				return tweets, err
			}
			tweets[i] = tweet
			prevStatusId = tweet.ID
		}
	}

	return tweets, nil
}

func (c *TwitterClient) Reply(msg string, tweetId int64) (*twitter.Tweet, error) {
	// Send a Tweet
	tweet, _, err := c.client.Statuses.Update(msg, &twitter.StatusUpdateParams{
		InReplyToStatusID:         tweetId,
		AutoPopulateReplyMetadata: Bool(true),
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", tweet)
	return tweet, nil
}

func (c *TwitterClient) CreateFavorite(tweetId int64) error {
	_, _, err := c.client.Favorites.Create(&twitter.FavoriteCreateParams{ID: tweetId})
	return err
}

func (c *TwitterClient) IsFollowing(screenName string) bool {
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

func SplitLongText(text string, maxLength int) []string {
	strLength := utf8.RuneCountInString(text)
	if strLength <= maxLength {
		return []string{text}
	}

	// 行の途中で分割されないように、行ごとに文字数チェック
	lines := strings.Split(text, "\n")
	count, str := utf8.RuneCountInString(lines[0]), lines[0]
	res := make([]string, 0)
	for i := 1; i < len(lines); i++ {
		lineCount := utf8.RuneCountInString(lines[i])
		if count+lineCount <= maxLength {
			count += lineCount
			str = strings.Join([]string{str, lines[i]}, "\n")
		} else {
			res = append(res, str)
			count, str = lineCount, lines[i]
		}
	}
	res = append(res, str)
	return res
}
