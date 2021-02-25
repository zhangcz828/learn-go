package main

import (
	"fmt"
)

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	str := "ab"

	fmt.Println(allStrings(str))
}

func allStrings(str string) []string {

	resStr := []string{}

	for i := range str {
		if len(resStr) == 0 {
			resStr = append(resStr, string(str[i]))
			continue
		}

		length := len(resStr)

		for _, subStr := range resStr {
			for j := 0; j <= len(resStr[0]); j++ {
				resStr = append(resStr, insertToString(subStr, str[i], j))
			}
		}

		resStr = resStr[length:]
	}

	return resStr
}

func insertToString(str string, b byte, location int) string {
	// insert b to str on location
	res := []byte{}

	if location == len(str) {
		return str + string(b)
	}

	for k, v := range []byte(str) {

		if k < location {
			res = append(res, v)
		} else {
			res = append(res, b)
			res = append(res, []byte(str)[location:]...)
			break
		}
	}

	return string(res)
}
