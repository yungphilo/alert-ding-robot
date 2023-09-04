package main

import (
	"fmt"
	"log"
	"os"

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

	prometheusMetricValue, err := GetMetricValue(pomUrl, metric)
	value := prometheusMetricValue.Data.Result[0].Value[1]
	// metric := config.PrometheusInfo.Metric
	values := GetInterfaceToInt(value)
	threshold := config.PrometheusInfo.Threshold
	if values > threshold {
		fmt.Printf("指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values)
		thresholds := FormatFileSize(int64(threshold))
		mvalue := FormatFileSize(int64(values))
		alertmesage := "指标disk：" + metric + "\n超出阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：http://grafana.soap.com/d/3Ra1cWRSk/test?orgId=1 \n"
		fmt.Println(alertmesage)
		//'"指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values'
		err = SendDingtalkMessage(&config, alertmesage)
		if err != nil {
			log.Fatalf("Failed to send Dingtalk message: %v", err)
		}

		fmt.Println("Dingtalk message sent successfully!")
	} else {
		fmt.Printf("指标 %s未超出阈值：%d \n当前值为：%d", metric, threshold, values)
	}

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
