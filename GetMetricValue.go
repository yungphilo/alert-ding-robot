package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// 发送请求获取值

func GetMetricValue(pomUrl, expr string) (PromPodDisk, error) {
	var promPodDisk PromPodDisk

	url := pomUrl + url.QueryEscape(expr)
	//url := pomUrl + metric
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("无法发送HTTP请求：%s\n", err.Error())
		return promPodDisk, err
	}
	defer resp.Body.Close()
	// var prometheusMetricValue PrometheusMetricValue
	err = json.NewDecoder(resp.Body).Decode(&promPodDisk)
	if err != nil {
		fmt.Printf("无法解析http响应：%s\n", err.Error())
		return promPodDisk, err
	}
	return promPodDisk, err

}
