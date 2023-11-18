package main

import "fmt"

func main(){
	number:= 10

	addressNumber := &number

	fmt.Print(*addressNumber)
}