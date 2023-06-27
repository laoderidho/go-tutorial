package main

import "fmt"

func main(){
	conditional(90)
	caseSwicth(80)
	caseConditional(70)
}

func conditional(hasil_ujian int ){
	if hasil_ujian >=90 {
		fmt.Println("A")
	}else if hasil_ujian >=80 {
		fmt.Println("B")
	}else if hasil_ujian >=70 {
		fmt.Println("C")
	}else if hasil_ujian >=60 {
		fmt.Println("D")
	}else{
		fmt.Println("F")
	}
}

func caseSwicth(hasil_ujian int ){
	switch hasil_ujian {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70:
		fmt.Println("C")
	case 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}

func caseConditional(hasil_ujian int){
	switch {
	case hasil_ujian >= 90:
		fmt.Println("A")
	case hasil_ujian >= 80:
		fmt.Println("B")
	case hasil_ujian >= 70:
		fmt.Println("C")
	case hasil_ujian >= 60:
		fmt.Println("D")
	}
}