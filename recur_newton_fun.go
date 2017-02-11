

package main

import (
	"fmt"
	"math"
)
var z= float64(1);
var count=int(0);
func cal(x,z float64,count int)float64{
	if count>30 {
		return z }
	
	val:=cal(x,z,count+1)
	
	return cal(x,val-(math.Pow(val,2)-x)/2,count+1);


}	
	func Sqrt(x float64) float64 {
	return cal(x,1,0);
 
}

func main() {
	fmt.Println(Sqrt(2))
}
