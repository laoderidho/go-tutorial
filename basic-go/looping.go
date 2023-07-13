package main 

import "fmt"

func main(){
	// forLoop()
	// breakLoop()
	continueLoop()
}

func forLoop(){
	for i := 10; i>0; i--{
		fmt.Println("perulangan ke ", i)
	}
}

func breakLoop(){
	for i := 10; i>0; i--{
		// break ini menghentikan perulangan jika sesuatu dalam kondisi ini terpenuhi
		if i == 5{
			break
		}
		fmt.Println("perulangan ke ", i)
	}
}

func continueLoop(){
	for i:= 0; i<10; i++{
		// continue ini melewati dan tidak akan ditampilkan jika kondisi ini terpenuhi
		if i%2==0{
			continue;
		}
		fmt.Println("perulangan ke", i)
	}
}
