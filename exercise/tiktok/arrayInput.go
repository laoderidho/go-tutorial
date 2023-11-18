package main

import "fmt"

func main(){
	myArr:= []int{76, 80, 66, 80, 98, 76, 54, 87}
	getArrayInput(myArr)
}

func getArrayInput(myarr []int){


	dataTerkecil := 0 
	dataTerbesar := 0 
	conditionPertama:= false
	resultSum:=0

	for _, data:= range myarr{
		if data > dataTerbesar{
			dataTerbesar = data
		}
		if !conditionPertama {
			dataTerkecil = dataTerbesar
			conditionPertama = true
		}
		if data < dataTerkecil{
			dataTerkecil = data
		}
		resultSum+= data
	}

	fmt.Println("data terbesar adalah ", dataTerbesar)
	fmt.Println("data terkecil adalah ", dataTerkecil)
	fmt.Println("jumlah keseluruhan data", resultSum)
	fmt.Println("jumlah rata rata", float32(resultSum/len(myarr)))
}