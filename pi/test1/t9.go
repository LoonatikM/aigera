package main

import "fmt"

func main() {
	sl := "substring"
	sl1 := []rune(sl)
	fmt.Println(string(sl1[2:6]))
}
