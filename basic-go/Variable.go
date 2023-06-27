package main

import "fmt"

func main(){
	fmt.Println("Hello Ridho")
	variable()
	fmt.Println(getSum(1,2));
	fmt.Println(getDataMhs("Ridho", 20, "123456789"))
	kewwordNew()
}

func variable(){
	// deklarasi var memuat sensitif tipe data
	var name string = "Ridho"

	// sementara deklarasi tanpa var tidak memuat sensitif tipe data
	age:= 20;

	fmt.Println(name, age)
}

func getSum(a int, b int) int {
	return a+b
}

// setiap parameter di golang harus memilkiki tipe data
func getDataMhs(mhs string, age int, nim string) (string, int, string){
	return mhs, age, nim
}

func kewwordNew(){
	name:= new(string)

	fmt.Println(name)
	fmt.Println(*name)
}
