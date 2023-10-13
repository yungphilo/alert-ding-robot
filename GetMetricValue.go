package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// 发送请求获取值

func GetMetricValue(client *http.Client, pomUrl, expr string) (PromPodDisk, error) {
	var promPodDisk PromPodDisk

	url := pomUrl + "/api/v1/query?query=" + url.QueryEscape(expr)
	//url := pomUrl + metric
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return promPodDisk, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("无法发送HTTP请求：%s\n", err.Error())
		return promPodDisk, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&promPodDisk)
	if err != nil {
		fmt.Printf("无法解析http响应：%s\n", err.Error())
		return promPodDisk, err
	}
	return promPodDisk, err

}
