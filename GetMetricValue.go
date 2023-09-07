package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 发送请求获取值

func GetMetricValue(pomUrl, metric string) (PromPodDisk, error) {
	var promPodDisk PromPodDisk

	url := pomUrl + metric + "%7Bimage!%3D%22%22%2Ccontainer%3D%7E%22.*order-.*%22%7D%2Fcontainer_spec_memory_limit_bytes%7Bimage!%3D%22%22%2Ccontainer%3D%7E%22.*test%22%7D%20*%20100"
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
