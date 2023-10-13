package main

import (
	"log"
	"strconv"
)

func compareInt(value interface{}, threshold int, metric string, podName string, nameSpace string, atmobiles []string, factor string, grafanaurl string) {
	values := GetInterfaceToInt(value)
	switch {
	case factor == "above":
		if values > threshold {

			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)

			alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "%" + "\n当前值为：" + mvalue + "%" + "\n" + "详情查看：" + grafanaurl
			log.Println(alertmesage)
			err := SendDingtalkMessage(&config, alertmesage, atmobiles)
			if err != nil {
				log.Fatalf("Failed to send Dingtalk message: %v", err)
			}
			log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
		} else {
			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)
			log.Printf("Pod %s指标 %s未超出阈值：%s \n当前值为：%s\n", podName, metric, thresholds, mvalue)
		}
	case factor == "below":
		if values < threshold {

			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)

			alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n低于阈值：" + thresholds + "%" + "\n当前值为：" + mvalue + "%" + "\n" + "详情查看：" + grafanaurl
			log.Println(alertmesage)
			err := SendDingtalkMessage(&config, alertmesage, atmobiles)
			if err != nil {
				log.Fatalf("Failed to send Dingtalk message: %v", err)
			}
			log.Printf("Dingtalk message sent successfully! @%s", atmobiles)
		} else {
			mvalue := strconv.Itoa(values)
			thresholds := strconv.Itoa(threshold)
			log.Printf("Pod %s指标 %s未超出阈值：%s \n当前值为：%s\n", podName, metric, thresholds, mvalue)
		}
	}

}
func compareFloat(value interface{}, threshold int, metric string, podName string, nameSpace string, atmobiles []string, grafanaurl string) {
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
func compareByte(value interface{}, threshold int, metric string, podName string, nameSpace string, atmobiles []string, grafanaurl string) {
	values := GetInterfaceToInt(value)
	if values > threshold {

		thresholds := FormatFileSize(int64(threshold))
		mvalue := FormatFileSize(int64(values))

		alertmesage := "pod disk 使用告警\n" + "指标pod disk：" + metric + "\nPod Name：" + podName + "\nNameSpace：" + nameSpace + "\n超出阈值：" + thresholds + "%" + "\n当前值为：" + mvalue + "%" + "\n" + "详情查看：" + grafanaurl
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
func comparePer(value interface{}, threshold int, metric string, podName string, nameSpace string, atmobiles []string, grafanaurl string) {
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

func listcompare(logstore []*string, job []string) ([]string, error) {
	diff := make([]string, 0)
	inter := SliceInterStr(logstore, job)
	str1 := make(map[string]int)
	for _, v := range inter {
		str1[v]++
	}
	for i := 0; i < len(logstore); i++ {
		times, ok := str1[*logstore[i]]
		if !ok || times == 0 {
			diff = append(diff, *logstore[i])
		}
		str1[*logstore[i]]++
	}

	return diff, nil
}

func SliceInterStr(slice1 []*string, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[*v]++
	}
	for _, v := range slice2 {
		times, ok := m[v]
		if ok && times > 0 {
			nn = append(nn, v)
			m[v]--
		}
	}
	return nn
}
