package main

import (
	"math"
	"fmt"
)

func main(){
	math:=50
	bIndo:=80
	english:=100
	science:=60

	fmt.Println(getPredicate(math, bIndo, english, science))
}

func getPredicate(grade ...int)string{
	var total int

	for _, num:= range(grade){
		total+=num
	}

	average:= float64(total) / float64(len(grade))

	sum:= int(math.Round(average))
	
	switch{
	case sum ==100:
		return "Sempurna"
	case sum > 90:
		return "Sangat Baik"
	case sum > 80:
		return "Baik"
	case sum > 70:
		return "Cukup"
	case sum>60:
		return "Kurang"
	default:
		return "tidak lulus"
	}
}