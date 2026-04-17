package main 
import "fmt"

func main(){
	sl:= make([]int, 0,5)
	sl= append(sl,2,3,6,7,8)
	fmt.Println(sl)
	sl=append(sl[:4],sl[4+1:]...)
	fmt.Println(sl)


}
