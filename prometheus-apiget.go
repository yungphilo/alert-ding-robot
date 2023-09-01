package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PrometheusInfo struct {
		URL       string `yaml:"url"`
		Metric    string `yaml:"metric"`
		Threshold int    `yaml:"threshold"`
	} `yaml:"prometheus"`
}

func main() {
	// 读取配置文件
	config, err := readConfig("config-a.yaml")
	if err != nil {
		fmt.Printf("无法读取配置文件：%s\n", err.Error())
		return
	}

	// 构建URL
	url := config.PrometheusInfo.URL + config.PrometheusInfo.Metric

	// 发送HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("无法发送HTTP请求：%s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	// 读取响应
	var prometheusMetricValue PrometheusMetricValue
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("无法读取HTTP响应：%s\n", err.Error())
		return prometheusMetricValue, err.Error()
	}

	// 打印结果
	fmt.Println(string(body))

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
