package main

import "fmt"

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePoint(&n)
	fmt.Println(a)
}

func UltimatePoint(n ***int) {
	***n = 1
}
