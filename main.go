package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	// 读取配置文件
	config, err := readConfig("config-text.yaml")
	if err != nil {
		fmt.Printf("无法读取配置文件：%s\n", err.Error())
		return
	}

	// // 构建URL

	pomUrl := config.PrometheusInfo.URL
	metric := config.PrometheusInfo.Metric

	//判断告警

	prometheusMetricValue, err := getMetricValue(pomUrl, metric)
	value := prometheusMetricValue.Data.Result[0].Value[1]
	// metric := config.PrometheusInfo.Metric
	values := GetInterfaceToInt(value)
	threshold := config.PrometheusInfo.Threshold
	if values > threshold {
		fmt.Printf("指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values)
		alertmesage := "指标disk：" + metric + "\n超出阈值：" + strconv.Itoa(threshold) + "\n当前值为：" + strconv.Itoa(values) + "\n" + "详情查看：http://grafana.soap.com/d/3Ra1cWRSk/test?orgId=1 \n"
		fmt.Println(alertmesage)
		//'"指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values'
		err = sendDingtalkMessage(&config, alertmesage)
		if err != nil {
			log.Fatalf("Failed to send Dingtalk message: %v", err)
		}

		fmt.Println("Dingtalk message sent successfully!")
	} else {
		fmt.Printf("指标 %s未超出阈值：%d \n当前值为：%d", metric, threshold, values)
	}

}

// 发送钉钉
func sendDingtalkMessage(config *Config, alertmesage string) error {
	payload := fmt.Sprintf(`{
    "msgtype": "%s",
    "%s": {
      "content": "%s"
    },
    "at": {
      "atMobiles": %s,
      "atUserIds": %s,
      "isAtAll": %t
    }
  }`, config.Message.MsgType, config.Message.MsgType, alertmesage, arrayToJSON(config.Message.At.AtMobiles), arrayToJSON(config.Message.At.AtUserIds), config.Message.At.IsAtAll)

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

// 发送请求获取值
func getMetricValue(pomUrl, metric string) (PrometheusMetricValue, error) {
	var prometheusMetricValue PrometheusMetricValue

	url := pomUrl + metric
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("无法发送HTTP请求：%s\n", err.Error())
		return prometheusMetricValue, err
	}
	defer resp.Body.Close()
	// var prometheusMetricValue PrometheusMetricValue
	err = json.NewDecoder(resp.Body).Decode(&prometheusMetricValue)
	if err != nil {
		fmt.Printf("无法解析http响应：%s\n", err.Error())
		return prometheusMetricValue, err
	}
	return prometheusMetricValue, err

}

func readConfig(filename string) (Config, error) {
	var config Config

	// 读取文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// 解析YAML
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
