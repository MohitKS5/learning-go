package main 

import fmt

var (
	tobe bool =false;
	Maxint uint64=1<<64-1;
	uninitialised int;
	  )

func main() {
	var a,b,c =true,23,"yeah";
	//equivalent statements
	var i int =1;
	//and
	i:=1;
	i,j,k:=1,2,3//also allowed

	fmt.Printf("type : %T value= %v %v",tobe,tobe,uninitialised);   //%v for value %T for datatype %q for string
}