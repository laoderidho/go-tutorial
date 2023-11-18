package main

import "fmt"

func main(){
	studentName := "Ridho Fahreza"
	studentAge:= 12
	grade := 90

	fmt.Println(returnStruct(studentName, studentAge))
	fmt.Println(embededStruct(grade))
}

type person struct{
	name string
	age int
}

func returnStruct(myName string, myAge int)person{

	var myPerson person
	
	myPerson.name = myName
	myPerson.age = myAge

	return myPerson
}

// embeded struct sistemnya sama dengan inheritance atau ovverriding yaitu menimpa struct 1 ke dalam struct 2
// dan juga sebaliknya 

type grade struct{
	// person di bawah di ambil dari person struct di atas
	person
	studentGrade int
}

func embededStruct(myGrade int)grade{

	var getGrade grade

	getGrade.name = "Ridho"

	getGrade.age= 12

	getGrade.studentGrade = myGrade

	return getGrade
}