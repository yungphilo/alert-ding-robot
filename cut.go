package main

import (
	"strings"
)

func Cutlast(name string) (cutname string) {
	svcname := strings.Split(name, "-")
	if len(svcname) > 1 {
		svcname = svcname[:len(svcname)-1]
	}
	cutname = strings.Join(svcname, "-")
	return cutname
}
