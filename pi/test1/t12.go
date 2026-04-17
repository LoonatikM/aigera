package main 
import ("fmt"
"math/rand/v2"
)

func main(){
	sl := make([]int, 5)
	for i:= range sl{
		sl[i]= rand.IntN(100)
	}
	fmt.Println(sl)

	sl[2] = sl[len(sl)-1]
	sl = sl[:len(sl)-1]
	fmt.Println(sl)

}