package main

import "fmt"

func main(){
	myAge:= 4
	myHeigth:= 85
	priceTicket:= validationWahana(myHeigth, myAge)
	fmt.Println(priceTicket)
}

func validationWahana(height int, age int) int {

	if age>12 {
		return 100000
	} else if age==12 || height > 160 {
		return 60000
	} else if age>=10 || height > 150{
		return 40000
	} else if age>=8 || height > 135{
		return 25000
	} else if age>=5 || height > 120{
		return 15000
	}else{
		return -1
	}
}