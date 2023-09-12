package main

import (
	"encoding/json"
	"strconv"
)

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

// interface to float
func GetInterfaceToFloat(t1 interface{}) float64 {
	var t2 float64
	switch t1.(type) {
	case uint:
		t2 = float64(t1.(uint))
		break
	case int8:
		t2 = float64(t1.(int8))
		break
	case uint8:
		t2 = float64(t1.(uint8))
		break
	case int16:
		t2 = float64(t1.(int16))
		break
	case uint16:
		t2 = float64(t1.(uint16))
		break
	case int32:
		t2 = float64(t1.(int32))
		break
	case uint32:
		t2 = float64(t1.(uint32))
		break
	case int64:
		t2 = float64(t1.(int64))
		break
	case uint64:
		t2 = float64(t1.(uint64))
		break
	case float32:
		t2 = float64(t1.(float32))
		break
	// case float64:
	// 	t2 = float64(t1.(float64))
	// 	break
	case string:
		t2, _ = strconv.ParseFloat(t1.(string), 64)
		if t2 == 0 && len(t1.(string)) > 0 {
			f, _ := strconv.ParseFloat(t1.(string), 64)
			t2 = float64(f)
		}
		break
	case nil:
		t2 = 0
		break
	case json.Number:
		t3, _ := t1.(json.Number).Int64()
		t2 = float64(t3)
		break
	default:
		t2 = t1.(float64)
		break
	}
	return t2
}
