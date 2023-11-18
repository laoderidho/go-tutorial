package main

import "fmt"

func main(){
	mygender:= "female"
	myHeigth:= 165

	fmt.Println(Bmi(mygender, myHeigth))
}

func Bmi(gender string, height int) float64{
	myHeigth:= float64(height)

	if gender == "male"{
		return (myHeigth - 100) - ((myHeigth - 100) * 0.1)
	}else{
		return (myHeigth - 100) -((myHeigth - 100)* 0.15)
	}
}