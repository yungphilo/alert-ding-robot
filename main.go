package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

func main() {
	//日志配置
	f, err := os.OpenFile("logs/service.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()
	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 读取配置文件
	config, err := readConfig("config-text.yaml")
	if err != nil {
		fmt.Printf("无法读取配置文件：%s\n", err.Error())
		return
	}

	// // 构建URL

	pomUrl := config.PrometheusInfo.URL
	metric := config.PrometheusInfo.Metric

	for {
		promPodDisk, err := GetMetricValue(pomUrl, metric)
		//fmt.Println(promPodDisk)
		atalerts := config.Atalerts
		fmt.Println(atalerts)
		for i := 0; i < len(promPodDisk.Data.Result); i++ {
			value := promPodDisk.Data.Result[i].Value[1]
			podName := promPodDisk.Data.Result[i].Metric.PodName
			nameSpace := promPodDisk.Data.Result[i].Metric.Namespace
			grafanaurl := config.PrometheusInfo.Grafana
			deployment := promPodDisk.Data.Result[i].Metric.Container
			service := Cutlast(deployment)
			atmobiles := FindMobiles(service, atalerts)
			values := GetInterfaceToInt(value)
			threshold := config.PrometheusInfo.Threshold * 3
			//log.Println(deployment)
			if values > threshold {
				//fmt.Printf("指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values)
				thresholds := FormatFileSize(int64(threshold))
				mvalue := FormatFileSize(int64(values))
				alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：" + grafanaurl
				log.Println(alertmesage)
				err = SendDingtalkMessage(&config, alertmesage, atmobiles)
				if err != nil {
					log.Fatalf("Failed to send Dingtalk message: %v", err)
				}
				log.Println("Dingtalk message sent successfully!")
			} else {
				log.Printf("Pod %s指标 %s未超出阈值：%s \n当前值为：%s\n", podName, metric, FormatFileSize(int64(threshold)), FormatFileSize(int64(values)))
			}
		}

		time.Sleep(5 * time.Minute)
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
