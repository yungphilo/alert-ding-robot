package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	} else {
		fmt.Printf("指标 %s未超出阈值：%d \n当前值为：%d", metric, threshold, values)
	}

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
<<<<<<< HEAD
	// var prometheusMetricValue PrometheusMetricValue
	err = json.NewDecoder(resp.Body).Decode(&prometheusMetricValue)
	if err != nil {
		fmt.Printf("无法解析http响应：%s\n", err.Error())
		return prometheusMetricValue, err
=======

	// 读取响应

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("无法读取HTTP响应：%s\n", err.Error())
		return
>>>>>>> 88a5f1c89f02f87f1a0693a2ad9a381f1d6fd2c6
	}
	return prometheusMetricValue, err

<<<<<<< HEAD
}
=======
	// 打印结果
	fmt.Println(string(body))
	//判断告警
	var prometheusMetricValue PrometheusMetricValue
	prometheusMetricValue, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	value := prometheusMetricValue.Data.Result.Value[2]
	values, err := strconv.Atoi(value)
	threshold := config.PrometheusInfo.Threshold
	if values > threshold {
		fmt.Printf("mtric %s超出阈值%s \n 当前值为%s", value, threshold, value)

	}
>>>>>>> 88a5f1c89f02f87f1a0693a2ad9a381f1d6fd2c6

// 将interface转成int
func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		if t2 == 0 && len(t1.(string)) > 0 {
			f, _ := strconv.ParseFloat(t1.(string), 64)
			t2 = int(f)
		}
		break
	case nil:
		t2 = 0
		break
	case json.Number:
		t3, _ := t1.(json.Number).Int64()
		t2 = int(t3)
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
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
