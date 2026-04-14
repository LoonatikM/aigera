package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	l := 0
	for i := 0; i < len(s); i++ {
		l = l + 1
	}
	return l
}
