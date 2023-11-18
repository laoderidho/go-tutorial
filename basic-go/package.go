package main

import (
	"strings"
	"fmt"
)

func main(){

	// mengubah huruf kapital menjadi huruf kecil
	fmt.Println("mengubah huruf HEBAT menjadi " + strings.ToLower("HEBAT"))

	// mengubah huruf kecil ke kapital 
	fmt.Println(strings.ToUpper("hello"))

	// mengubah indeks string
	// argument 1: nama, argumen 2: huruf yang ingin diganti, 3: huruf pengganti, 4: berapa yang ingin diganti
	// jika ingin menganti semua huruf a maka argument ke empat jadikan -1
	fmt.Println(strings.Replace("banana", "a", "i", 1))
	fmt.Printf(strings.Replace("banana", "a", "i", -1))
}