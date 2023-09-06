package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 发送请求获取值
func GetMetricValue(pomUrl, metric rune) (PromPodDisk, error) {
	var prometheusMetricValue PromPodDisk

	url := pomUrl + string(metric) + "{pod%3D%7E%22eastbuy-xxl-job-admin-test.*%22%2C%20device%3D%22%2Fdev%2Fvdb%22%2C%20container%3D%22eastbuy-xxl-job-admin-test%22}"
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
