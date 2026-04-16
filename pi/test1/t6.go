package main

import "fmt"

func main() {
	var b bool = true
	p := &b
	fmt.Println(*p)
	*p = false
	fmt.Println(b)
}
