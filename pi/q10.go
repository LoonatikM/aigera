package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2(""))
}

func BasicAtoi2(s string) int {
	res := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
		}
		if string(c) == " " {
			return 0
		}
	}

	return res
}
