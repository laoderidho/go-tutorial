package main 

import "fmt"

func main(){
	fmt.Println(myFunction(1,2,3,1,4,1,3,1,31,21));
	
	
	num1:= func(myNum int)int{
		return myNum +1
	}

	fmt.Println(ParameterFunc(num1, 12))

	name:="Ridho"
	fmt.Println(invoked(name))

	fmt.Println(factorial(5))
}

// parameter harus mempunyai tipe data 
// 
func myFunction(f1 ...int) int{
	sum := 0;
	for _, value := range f1{
		sum += value;
	}
	return sum;
}

func ParameterFunc(f func(x int) int, number int) int{
	return f(number)
}


func num2(myNum int) int{
	return myNum + 2;
}

// IIFE function adalah sistem yang membuat function yang memiliki parameter dan argument yang di deklarasikan di functionnya

func invoked(testing string)string{
	
	invokedFunc:= func(name string)string{
		return "Hello " + name
	}(testing)

	return invokedFunc
}

// recursive function
// dari fungsi ini akan berjalan terus karena memanggil fungsinya sendiri 
// ini membuat fungsi ini harus dihentikan jika kondisi terpenuhi
func factorial(params int) int{

	if params == 0{
		return 1
	}else{
		return params * factorial(params -1)
	}
}
