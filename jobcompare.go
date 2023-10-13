package main

import (
	"log"
	"strconv"
)

func jobcompareInt(value interface{}, threshold int, alertname string, atmobiles []string, factor string, grafanaurl string) {
	values := GetInterfaceToInt(value)
	switch {
	case factor == "above":
		if values > threshold {

			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)

			alertmesage := "监控告警\n" + "告警名称：" + alertname + "\n超出阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：" + grafanaurl
			log.Println(alertmesage)
			err := SendDingtalkMessage(&config, alertmesage, atmobiles)
			if err != nil {
				log.Fatalf("Failed to send Dingtalk message: %v", err)
			}
			log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
		} else {
			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)
			log.Printf("告警名称： %s未超出阈值：%s \n当前值为：%s\n", alertname, thresholds, mvalue)
		}
	case factor == "below":
		if values < threshold {

			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)

			alertmesage := "监控告警\n" + "告警名称：" + alertname + "\n低于阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：" + grafanaurl
			log.Println(alertmesage)
			err := SendDingtalkMessage(&config, alertmesage, atmobiles)
			if err != nil {
				log.Fatalf("Failed to send Dingtalk message: %v", err)
			}
			log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
		} else {
			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)
			log.Printf("告警名称： %s未超出阈值：%s \n当前值为：%s\n", alertname, thresholds, mvalue)
		}
	}

}
func jobcompareFloat(value interface{}, threshold int, metric string, podName string, nameSpace string, atmobiles []string, grafanaurl string) {
	values := GetInterfaceToFloat(value)
	thresholds := float64(threshold)
	if values > thresholds {

		thresholds := strconv.FormatFloat(thresholds, 'f', 3, 64)
		mvalue := strconv.FormatFloat(values, 'f', 3, 64)

		alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：" + grafanaurl
		log.Println(alertmesage)
		err := SendDingtalkMessage(&config, alertmesage, atmobiles)
		if err != nil {
			log.Fatalf("Failed to send Dingtalk message: %v", err)
		}
		log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
	} else {
		mvalue := strconv.FormatFloat(values, 'f', 3, 64)
		thresholds := strconv.FormatFloat(thresholds, 'f', 3, 64)
		log.Printf("Pod %s指标 %s未超出阈值：%s \n当前值为：%s\n", podName, metric, thresholds, mvalue)
	}
}
func jobcompareByte(value interface{}, threshold int, metric string, podName string, nameSpace string, atmobiles []string, grafanaurl string) {
	values := GetInterfaceToInt(value)
	if values > threshold {

		thresholds := FormatFileSize(int64(threshold))
		mvalue := FormatFileSize(int64(values))

		alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "\n当前值为：" + mvalue + "\n" + "详情查看：" + grafanaurl
		log.Println(alertmesage)
		err := SendDingtalkMessage(&config, alertmesage, atmobiles)
		if err != nil {
			log.Fatalf("Failed to send Dingtalk message: %v", err)
		}
		log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
	} else {
		mvalue := FormatFileSize(int64(values))
		thresholds := FormatFileSize(int64(threshold))
		log.Printf("Pod %s指标 %s未超出阈值：%s \n当前值为：%s\n", podName, metric, thresholds, mvalue)
	}
}
func jobcomparePer(value interface{}, threshold int, metric string, podName string, nameSpace string, atmobiles []string, grafanaurl string) {
	values := GetInterfaceToFloat(value)
	thresholds := float64(threshold)
	if values > thresholds {

		thresholds := strconv.FormatFloat(thresholds, 'f', 3, 64)
		mvalue := strconv.FormatFloat(values, 'f', 3, 64)

		alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "%" + "\n当前值为：" + mvalue + "%" + "\n" + "详情查看：" + grafanaurl
		log.Println(alertmesage)
		err := SendDingtalkMessage(&config, alertmesage, atmobiles)
		if err != nil {
			log.Fatalf("Failed to send Dingtalk message: %v", err)
		}
		log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
	} else {
		mvalue := strconv.FormatFloat(values, 'f', 3, 64)
		thresholds := strconv.FormatFloat(thresholds, 'f', 3, 64)
		log.Printf("Pod %s指标 %s未超出阈值：%s \n当前值为：%s\n", podName, metric, thresholds, mvalue)
	}

}
