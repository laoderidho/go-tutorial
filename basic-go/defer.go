package main

import "fmt"

func main(){
	myDefer()
}

func myDefer(){
	defer fmt.Println("this print last because this defer")
	fmt.Println("this print first")
	fmt.Println("this print Second")
}