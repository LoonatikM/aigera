package main

import "fmt"

func main() {
	// 1
	s := []int{1, 6, 9, 8, 5}
	fmt.Printf("len: %v \nCap: %v\n", len(s), cap(s))
	// 2
	s = append(s, 3, 2, 6)
	fmt.Printf("SliceNew:%v \nlenNew: %v \nCapNew: %v\n", s, len(s), cap(s))
	// 3
	d := s[1 : len(s)-1]
	fmt.Println("Cut Slice:", d)
	// 4
	sl := "GO"
	r := []rune(sl)
	fmt.Println(string(r[0]))
	// 5
	q := 5
	p := &q
	fmt.Println(*p)
	*p = 1213
	fmt.Println(*p)
}
