package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 发送请求获取值
func GetMetricValue(pomUrl, metric string) (PrometheusMetricValue, error) {
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
