package main

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
		URL     string `yaml:"url"`
		Metrics []struct {
			Metric    string `yaml:"metric"`
			Threshold int    `yaml:"threshold"`
			Grafana   string `yaml:"grafanaurl"`
			Expr      string `yaml:"expr"`
			Type      string `yaml:"type"`
		} `yaml:"metrics"`
		Window int `yaml:"window"`
	} `yaml:"prometheus"`
	DingtalkWebhook string `yaml:"dingtalk_webhook"`
	Secret          string `yaml:"secret"`
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
	Atalerts map[string][]string `yaml:"atalerts"`
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
