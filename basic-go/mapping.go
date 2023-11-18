package main

import "fmt"

func main(){
	fmt.Println(dataMap())
}

// data map di definisikan dengan cara make (map[namakey]namavalue)
func dataMap()map[string]string {
	data:= map[string]string{
		"testing1": "myTest",
		"testing2": "my Test Go",
	}

	// read map

	myScore := map[int]string{
		100: "A",
		90: "B",
		80: "C",
		70: "D",
		60: "F",
	}

	// ambil key untuk melihat isi value
	fmt.Println("key nya adalah 100 dan valuenya adalah " + myScore[100])

	// update key and value

	myNameAndGrade:= map[string]int{
		"Ridho": 80,
		"Fahreza": 67,
	}

	myNameAndGrade["Ridho"] = 100

	fmt.Println(myNameAndGrade["Ridho"])

	// data["testing"] = "tes1"
	// data["testing2"] ="test2"
	
	
	return data
}