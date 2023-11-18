package main

import (
	"strings"
	"fmt"
)

func main(){
	myWord:= "Zebra Zig Zag";
	fmt.Println(countingLetter(myWord))
}

func countingLetter(params string)int{
	cadel := []string{"r", "s", "t", "z"}
	getParams := strings.ToLower(params) 
	countCadel:= 0;

	for i:=0; i< len(getParams); i++{
		for j:=0; j<len(cadel); j++{
			if string(getParams[i]) == cadel[j]{
				countCadel++;
			}
		}
	}

	return countCadel
}