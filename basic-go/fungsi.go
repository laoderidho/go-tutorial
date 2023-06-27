package main

import(
	"fmt"
	"strings"
)


func main(){
	var names = []string{"john", "doe"}
	printMessage("halo, ", names)

	var numberData = calculate(1,3,2,3,1,1,44,1,131,121,31121,)
	fmt.Println(numberData)
}

func printMessage(message string, arr []string){
	var joinstring = strings.Join(arr, " ")
	fmt.Println(message, joinstring)
}

func calculate (number ...int)float64{
	var total int= 0
	
	for _, number:= range number{
		total+= number
	}

	var avg = float64(total) / float64(len(number))

	return avg
}