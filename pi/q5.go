package main

import (
	"fmt"
)

func main() {
	PrintStr("Hello World!")
}

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Print("\n")
}

// s[i]

// for (перемен); (условия  переменая < len(s)); (i++){

// }
