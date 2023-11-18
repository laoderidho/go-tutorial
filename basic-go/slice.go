package main 

import "fmt"

func main() {
	val := []string{"ridho", "fahreza", "La Ode"}
	sliceArray(val)
	
	variable:= []int{1,2,3,4,5,6,6,7}
	value:= 2
	fmt.Println(getSliceArray(variable, value))

	stringVar:= []string{"ridho, Fahreza"}
	valueString:= "la Ode"

	fmt.Println(getAppend(stringVar, valueString))
}

func sliceArray(val []string){
	for i := 0; i < len(val); i++ {
		fmt.Print(val[i])
	}
}

func getSliceArray(variable []int, value int) []int {
	for i := 1; i < len(variable); i++ {
		if variable[i-1] == value {
			variable[i-1] -= value
		}
	}

	return variable
}

// append array 
func getAppend(data[]string, value string) []string{
	myData:= append(data, value)

	return myData
}