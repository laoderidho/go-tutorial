package main

import (
	"fmt"
	"test-module/calculator"
	"test-module/person"
)

func main(){
	number1:= 12
	number2:= 21

	myName:= "Ridho Fahreza"
	fmt.Println(calculator.Plus(number1, number2))
	
	getDataPerson := person.MyPerson{
		Name: myName,
		Age: 22,
		Grade: 90,
	}

	fmt.Println(getDataPerson.SayHello())
}