package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

func main() {
	//日志配置
	f, err := os.OpenFile("logs/service.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("no logs directory，will create path ")
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("logs create")
	}
	defer func() {
		f.Close()
	}()
	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 读取配置文件
	config, err := readConfig("conf/config.yaml")
	if err != nil {
		fmt.Printf("无法读取配置文件：%s\n", err.Error())
		return
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	// // 构建URL

	pomUrl := config.PrometheusInfo.URL

	for {
		for j := 0; j < len(config.PrometheusInfo.Metrics); j++ {
			metric := config.PrometheusInfo.Metrics[j].Metric
			expr := config.PrometheusInfo.Metrics[j].Expr
			grafanaurl := config.PrometheusInfo.Metrics[j].Grafana
			threshold := config.PrometheusInfo.Metrics[j].Threshold
			types := config.PrometheusInfo.Metrics[j].Type
			promPodDisk, err := GetMetricValue(&client, pomUrl, expr)
			atalerts := config.Atalerts
			fmt.Println(atalerts)
			for i := 0; i < len(promPodDisk.Data.Result); i++ {
				value := promPodDisk.Data.Result[i].Value[1]
				podName := promPodDisk.Data.Result[i].Metric.PodName
				nameSpace := promPodDisk.Data.Result[i].Metric.Namespace

				deployment := promPodDisk.Data.Result[i].Metric.Container
				//deployment名字为服务名字+“-” +环境变量，去掉“-”及后面的环境参数
				service := Cutlast(deployment)
				atmobiles := FindMobiles(service, atalerts)
				switch {
				case types == "int":

				}
				values := GetInterfaceToFloat(value)

				//log.Println(deployment)
				if values > threshold {
					//fmt.Printf("指标 %s超出阈值：%d \n当前值为：%d", metric, threshold, values)
					//thresholds := FormatFileSize(int64(threshold))
					//mvalue := FormatFileSize(int64(values))
					thresholds := strconv.FormatFloat(threshold, 'f', 3, 64)
					mvalue := strconv.FormatFloat(values, 'f', 3, 64)
					// thresholds := float64(threshold)
					// mvalue := float64(values)
					alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "%" + "\n当前值为：" + mvalue + "%" + "\n" + "详情查看：" + grafanaurl
					log.Println(alertmesage)
					err = SendDingtalkMessage(&config, alertmesage, atmobiles)
					if err != nil {
						log.Fatalf("Failed to send Dingtalk message: %v", err)
					}
					log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
				} else {
					log.Printf("Pod %s指标 %s未超出阈值：%.2f%% \n当前值为：%.3f%%\n", podName, metric, threshold, values)
				}
			}
		}
		// times := config.PrometheusInfo.Window
		// log.Print(config.PrometheusInfo.Window)
		// tt := time.Duration(times.Minutes())
		tt := time.Duration(config.PrometheusInfo.Window) * time.Minute
		log.Println(tt)
		time.Sleep(time.Duration(config.PrometheusInfo.Window) * time.Minute)

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
