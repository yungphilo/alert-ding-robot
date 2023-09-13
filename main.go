package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

var config, readerr = readConfig("conf/config.yaml")

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
	// config, err := readConfig("conf/config.yaml")
	// if err != nil {
	// 	fmt.Printf("无法读取配置文件：%s\n", err.Error())
	// 	return
	if readerr != nil {
		log.Fatalf("无法读取配置文件：%s\n", readerr.Error())
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
			log.Println(err)
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
					compareInt(value, threshold, metric, podName, nameSpace, atmobiles, grafanaurl)
				case types == "float":
					compareFloat(value, threshold, metric, podName, nameSpace, atmobiles, grafanaurl)
				case types == "byte":
					compareByte(value, threshold, metric, podName, nameSpace, atmobiles, grafanaurl)
				case types == "per":
					comparePer(value, threshold, metric, podName, nameSpace, atmobiles, grafanaurl)

				}

				//log.Println(deployment)

			}
		}
		// times := config.PrometheusInfo.Window
		// log.Print(config.PrometheusInfo.Window)
		// tt := time.Duration(times.Minutes())
		// tt := time.Duration(config.PrometheusInfo.Window) * time.Minute
		// log.Println(tt)
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
