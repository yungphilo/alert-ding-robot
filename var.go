package main

import "time"

type PrometheusMetricValue struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name     string `json:"__name__,omitempty"`
				Instance string `json:"instance,omitempty"`
				Job      string `json:"job,omitempty"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// pod disk info data json
type PromPodDisk struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name          string `json:"__name__"`
				Container     string `json:"container"`
				Device        string `json:"device"`
				Endpoint      string `json:"endpoint"`
				ID            string `json:"id"`
				Image         string `json:"image"`
				Instance      string `json:"instance"`
				Job           string `json:"job"`
				ContainerName string `json:"name"`
				Namespace     string `json:"namespace"`
				Node          string `json:"node"`
				Pod           string `json:"pod"`
				PodName       string `json:"pod_name"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
type Config struct {
	PrometheusInfo struct {
		URL       string        `yaml:"url"`
		Metric    string        `yaml:"metric"`
		Threshold int           `yaml:"threshold"`
		Window    time.Duration `yaml:"window"`
		Minute    time.Duration `yaml:"minute"`
		Grafana   string        `yaml:"grafanaurl"`
	} `yaml:"prometheus"`
	DingtalkWebhook string `yaml:"dingtalk_webhook"`
	Message         struct {
		MsgType string `yaml:"msgtype"`
		Text    struct {
			Content string `yaml:"content"`
		} `yaml:"text"`
		At struct {
			AtMobiles []string `yaml:"atMobiles"`
			AtUserIds []string `yaml:"atUserIds"`
			IsAtAll   bool     `yaml:"isAtAll"`
		} `yaml:"at"`
	} `yaml:"dingmessage"`
}

// type Config struct {
// 	DingtalkWebhook string `yaml:"dingtalk_webhook"`
// 	Message         struct {
// 		MsgType string `yaml:"msgtype"`
// 		Text    struct {
// 			Content string `yaml:"content"`
// 		} `yaml:"text"`
// 		At struct {
// 			AtMobiles []string `yaml:"atMobiles"`
// 			AtUserIds []string `yaml:"atUserIds"`
// 			IsAtAll   bool     `yaml:"isAtAll"`
// 		} `yaml:"at"`
// 	} `yaml:"message"`
// }
