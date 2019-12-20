package twitter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func (c *TwitterClient) CreateCRCToken(crcToken string) string {
	mac := hmac.New(sha256.New, []byte(c.envConfig.ConsumerSecret))
	mac.Write([]byte(crcToken))
	return "sha256=" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
