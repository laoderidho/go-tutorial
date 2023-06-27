package main

import "fmt"

func main(){
	fmt.Println(getArray())
	
	mahasiswaData()
	lenArr()
}

func getArray()[3]string{
	var arr= [3] string{ "ridho", "fahreza", "ridho fahreza" }
	return arr
}

func lenArr(){
	var arr= [3] string{ "ridho", "fahreza", "ridho fahreza" }
	for i:=0; i<len(arr); i++{
		fmt.Println(arr[i])
		if(arr[i] == "ridho"){
			fmt.Println("Nama saya ridho")
		}else{
			fmt.Println("Nama saya bukan ridho")
		}
	}
}

// array of object 

type mahasiswa struct{
	name string
	NIM string
	age int
}

func mahasiswaData() [3]mahasiswa{
	getDataMahasiswa:= [3]mahasiswa{
		{name: "Ridho", NIM: "2081007", age: 21},
		{name: "Syafa", NIM: "2081021", age: 20},
		{name: "ridho Fahreza", NIM: "21321391", age: 20},
	}

	for _, mhs:= range getDataMahasiswa{
		fmt.Println("Name", mhs.name)
		fmt.Println("NIM", mhs.NIM)
		fmt.Println("Age", mhs.age)
	}

	return getDataMahasiswa
}