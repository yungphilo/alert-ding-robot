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

	//判断告警

	// prometheusMetricValue, err := GetMetricValue(pomUrl, metric)
	// value := prometheusMetricValue.Data.Result[0].Value[1]
	// // metric := config.PrometheusInfo.Metric
	// values := GetInterfaceToInt(value)
	// threshold := config.PrometheusInfo.Threshold
	// if values > threshold {
	// 	fmt.Printf("指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values)
	// 	//单位换算
	// 	thresholds := FormatFileSize(int64(threshold))
	// 	mvalue := FormatFileSize(int64(values))
	// 	alertmesage := "指标disk：" + metric + "\n超出阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：http://grafana.soap.com/d/3Ra1cWRSk/test?orgId=1 \n"
	// 	fmt.Println(alertmesage)
	// 	//'"指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values'
	// 	err = SendDingtalkMessage(&config, alertmesage)
	// 	if err != nil {
	// 		log.Fatalf("Failed to send Dingtalk message: %v", err)
	// 	}

	// 	fmt.Println("Dingtalk message sent successfully!")
	// } else {
	// 	fmt.Printf("指标 %s未超出阈值：%d \n当前值为：%d", metric, threshold, values)
	// }
	//

	for {
		promPodDisk, err := GetMetricValue(pomUrl, metric)
		//fmt.Println(promPodDisk)
		for i := 0; i < len(promPodDisk.Data.Result); i++ {
			value := promPodDisk.Data.Result[i].Value[1]
			podName := promPodDisk.Data.Result[0].Metric.PodName
			nameSpace := promPodDisk.Data.Result[0].Metric.Namespace
			grafanaurl := config.PrometheusInfo.Grafana
			//service := promPodDisk.Data.Result[0].Metric.Container
			values := GetInterfaceToInt(value)
			threshold := config.PrometheusInfo.Threshold * 3
			if values > threshold {
				//fmt.Printf("指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values)
				thresholds := FormatFileSize(int64(threshold))
				mvalue := FormatFileSize(int64(values))
				alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：" + grafanaurl
				log.Println(alertmesage)

				err = SendDingtalkMessage(&config, alertmesage)
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

// func toMap(filename string) (Config, error) {
// 	var user Config
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return user, err
// 	}
// 	err = yaml.Unmarshal(data, &user)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return user, nil
// }
