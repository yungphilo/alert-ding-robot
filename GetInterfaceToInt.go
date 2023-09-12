package main

import (
	"encoding/json"
	"strconv"
)

// 将interface转成int
func GetInterfaceToInt(t1 interface{}) int {
	var t2 int

	switch t1 := t1.(type) {
	case uint:
		t2 = int(t1)
	case int8:
		t2 = int(t1)
	case uint8:
		t2 = int(t1)
	case int16:
		t2 = int(t1)
	case uint16:
		t2 = int(t1)
	case int32:
		t2 = int(t1)
	case uint32:
		t2 = int(t1)
	case int64:
		t2 = int(t1)
	case uint64:
		t2 = int(t1)
	case float32:
		t2 = int(t1)
	case float64:
		t2 = int(t1)
	case string:
		t2, _ = strconv.Atoi(t1)
		if t2 == 0 && len(t1) > 0 {
			f, _ := strconv.ParseFloat(t1, 64)
			t2 = int(f)
		}
	case nil:
		t2 = 0
	case json.Number:
		t3, _ := t1.Int64()
		t2 = int(t3)
	default:
		t2 = t1.(int)
	}
	return t2
}

// interface to float
func GetInterfaceToFloat(t1 interface{}) float64 {
	var t2 float64
	switch t1 := t1.(type) {
	case uint:
		t2 = float64(t1)
	case int8:
		t2 = float64(t1)
	case uint8:
		t2 = float64(t1)
	case int16:
		t2 = float64(t1)
	case uint16:
		t2 = float64(t1)
	case int32:
		t2 = float64(t1)
	case uint32:
		t2 = float64(t1)
	case int64:
		t2 = float64(t1)
	case uint64:
		t2 = float64(t1)
	case float32:
		t2 = float64(t1)
	// case float64:
	// 	t2 = float64(t1.(float64))
	// 	break
	case string:
		t2, _ = strconv.ParseFloat(t1, 64)
		if t2 == 0 && len(t1) > 0 {
			f, _ := strconv.ParseFloat(t1, 64)
			t2 = float64(f)
		}
	case nil:
		t2 = 0
	case json.Number:
		t3, _ := t1.Int64()
		t2 = float64(t3)
	default:
		t2 = t1.(float64)
	}
	return t2
}
