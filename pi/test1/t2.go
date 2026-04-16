package main

import "fmt"

func main() {
	sl := []string{"Aigerim", "Ghsgdgs", "Cdjlfik", "dhfhgbv"}
	var sl1 string
	for i := 0; i < len(sl); i++ {
		sl1 += sl[i] + " "
	}

	fmt.Println(sl1)
}
