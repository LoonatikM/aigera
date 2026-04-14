package main

import (
	"fmt"
)

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a, "\n", b)
}

func UltimateDivMod(a *int, b *int) {
	temp := *a
	*a = *a / *b
	*b = temp % *b
}
