package main

	
import ("fmt" "math/rand")

func add(x int ,y int )int {
	reuturn x+y
}

func swap(x,y string )(string,string) {
	reuturn y,x
	
}
func main() {
	fmt.println(rand.Intn(10),add(1,2))
	fmt.println(swap("hell","bell "))
}