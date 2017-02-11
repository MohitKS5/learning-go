package main

import ("fmt" 
 "math")

func example(x, n , lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}


func main() {
	
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}

	fmt.Println(example(3,4,20));
	
}
