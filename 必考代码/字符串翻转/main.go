package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(reverseString(s))
}

func reverseString(str string) string {
	r := []rune(str) // 为什么不用byte
	l := len(str)
	for i := 0; i < l/2; i++ {
		fmt.Println(r)
		r[i], r[l-1-i] = r[l-1-i], r[i]
	}

	return string(r)
}
