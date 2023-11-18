package main

import "fmt"

func main(){
	ridhoAbsent:= 4
	ridhoExamp:= 70

	fmt.Println(isPass(ridhoExamp, ridhoAbsent))
}

func isPass (exam int, absent int) string{
	if exam >= 70 && absent < 5{
		return "pass"
	}else{
		return "Not Pass"
	}
}