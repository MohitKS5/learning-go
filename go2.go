package main

	
import ("fmt" "math/rand")

func add(x int ,y int )int {
	reuturn x+y
}
func main() {
	fmt.println(rand.Intn(10),add(1,2))
}