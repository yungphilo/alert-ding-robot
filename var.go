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
type Config struct {
	PrometheusInfo struct {
		URL       string `yaml:"url"`
		Metric    string `yaml:"metric"`
		Threshold int    `yaml:"threshold"`
	} `yaml:"prometheus"`
}
