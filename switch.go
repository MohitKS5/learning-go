package main 

import ("fmt"
		"time")

func main() {
	day:=time.Saturday;
	today:=time.Now().Weekday();
	switch(day){
	case today:fmt.Println("today")
	case today+1:fmt.Println("tomorrow")
	case today+2:fmt.Println("day after tomorrow")
    case today+3:fmt.Println("in two days")
	}


}