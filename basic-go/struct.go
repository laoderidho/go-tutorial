package main

import "fmt"

type mhs struct {
	name string
	age int
	nim string
}

func main(){
	fmt.Println(MyStruct())
	getAllStudent()
}


// fungsi dibawah ini mengembalikan objek
func MyStruct() mhs{
	var mhs1 mhs

	mhs1.name = "Rizky"
	mhs1.age = 23
	mhs1.nim = "G64170001"

	return mhs1
}

func getAllStudent(){

	var allStudent = []mhs{
		{name: "Ridho", age: 21, nim: "2081007"},
		{name: "Syafa", age: 21, nim: "2081002"},
		{name: "Salwa", age: 20, nim: "2018311"},
	}

	for _, student:= range allStudent{
		fmt.Println(student.name, "age", student.age, "NIM", student.nim)
	}
}