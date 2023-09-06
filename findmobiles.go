package main

import "fmt"

func FindMobiles(servername string, atalerts map[string][]string) (mobiles []string) {
	svcname, ok := atalerts[servername]
	if ok {
		fmt.Println(svcname)
		// multmumber := len(atalerts[servername])
		// for j := 0; j < multmumber; j++ {
		// 	phone := atalerts[servername][j]
		// }
		mobiles := atalerts[servername]
		return mobiles

	}
	return mobiles
}
