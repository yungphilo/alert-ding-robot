package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 发送钉钉
func SendDingtalkMessage(config *Config, alertmesage string, atmobiles []string) error {
	payload := fmt.Sprintf(`{
    "msgtype": "%s",
    "%s": {
      "content": "%s"
    },
    "at": {
      "atMobiles": %s,
      "isAtAll": %t
    }
  }`, config.Message.MsgType, config.Message.MsgType, alertmesage, arrayToJSON(atmobiles), config.Message.At.IsAtAll)
	//d := ding.Webhook{AccessToken: config.Token, Secret: config.Secret}
	//_ = d.SendMessageText(strings.NewReader(payload))
	//_ = d.SendDingtalkMessage(strings.NewReader(payload))
	if config.Secret == "" {
		// timestamp := time.Now().UnixNano()
		// time := strconv.Itoa(int(timestamp))
		dingUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + config.Token
		resp, err := http.Post(dingUrl, "application/json", strings.NewReader(payload))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to send Dingtalk message. StatusCode: %d", resp.StatusCode)
		}
	} else {
		timestamp := time.Now().UnixNano() / 1e6
		time := strconv.Itoa(int(timestamp))
		stringToSign := fmt.Sprintf("%s\n%s", time, config.Secret)
		signSha := hmac.New(sha256.New, []byte(config.Secret))
		signSha.Write([]byte(stringToSign))
		signData := signSha.Sum(nil)
		sign := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
		dingUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + config.Token + "&timestamp=" + time + "&sign=" + sign
		resp, err := http.Post(dingUrl, "application/json", strings.NewReader(payload))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to send Dingtalk message. StatusCode: %d", resp.StatusCode)
		}
	}

	return nil
}

func arrayToJSON(arr []string) string {
	str := `["` + strings.Join(arr, `","`) + `"]`
	return str
}
