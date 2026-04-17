package main 
import (
	"fmt"
	"math/rand/v2"
)

func main(){
	sl:= make([]int, 10)
	Filt(&sl)
	fmt.Println(sl)
	Filt(&sl)
}

func Filt(p *[]int) {
	for i:= range *p {
		(*p)[i]= rand.IntN(100)-50
		return *p
	}
	new:=0
	for j:=0; j<len(*p); j++{
	if (*p)[j]>=0 {
		new+=(*p)[i]
	}
}
}

