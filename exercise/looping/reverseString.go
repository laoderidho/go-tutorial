package main

import "fmt"

func main(){
	myString:= "Hello Ridho";
	fmt.Print(reverString(myString))
}

func reverString(myString string) string{

	reverString:= ""
	var outputReverse string
	
	for i := len(myString)-1; i >= 0; i-- {
		reverString+= string(myString[i])
	}

	for j:= 0; j<len(reverString); j++{
		if (j + 1< len(reverString) && reverString[j+1] == ' ') || reverString[j]==' '{
				outputReverse+=string(reverString[j])
		}else{
			outputReverse+=string(reverString[j]) + "_"
		}
	}

	return outputReverse[:len(outputReverse)-1];
}

