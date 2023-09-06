package main

import (
	"fmt"
	"net/http"
	"strings"
)

// 发送钉钉
func SendDingtalkMessage(config *Config, alertmesage string) error {
	payload := fmt.Sprintf(`{
    "msgtype": "%s",
    "%s": {
      "content": "%s"
    },
    "at": {
      "atMobiles": %s,
      "isAtAll": %t
    }
  }`, config.Message.MsgType, config.Message.MsgType, alertmesage, arrayToJSON(config.Message.At.AtMobiles), config.Message.At.IsAtAll)

	resp, err := http.Post(config.DingtalkWebhook, "application/json", strings.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send Dingtalk message. StatusCode: %d", resp.StatusCode)
	}

	return nil
}

func arrayToJSON(arr []string) string {
	str := `["` + strings.Join(arr, `","`) + `"]`
	return str
}
