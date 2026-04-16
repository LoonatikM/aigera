package main

import "fmt"

func main() {
	sl := []string{"goo", "gooo", "goooo"}
	sl1 := make([]string, 3, 6)
	copy(sl1, sl)
	fmt.Println(sl1)
}
