package main

import "fmt"

func main(){
	getMap()
	getresult:= "IV"
	fmt.Println(romantoInt(getresult))
}

func getMap(){
	var data = map[int]string{
		2081007: "Ridho",
		2081006: "stanlay",
		2081035: "Anwar",
	}

	for key, val := range data{
		fmt.Println(key, " \t", val)
	}
}

func romantoInt(varroman string) int {
	temp:= 0
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i:= 0; i<len(varroman); i++{
		if i< len(varroman)-1 && roman[string(varroman[i])] < roman[string(varroman[i+1])]{
			temp -= roman[string(varroman[i])]
		}else{
			temp += roman[string(varroman[i])]
		}
	}

	return temp
}