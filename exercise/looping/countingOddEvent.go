package main

import (
	"fmt"
	// "math"
)

func main(){
	myCount:= 5
	fmt.Print(countingEvent(myCount))
}


func countingEvent(params int)[]float64{

	var getParams float64 = 1
	lengthParams:= float64(params)

	var arr []float64
	
	for i:=getParams; i<=lengthParams; i += 0.5{
		arr = append(arr, i)
	}

	return arr
}