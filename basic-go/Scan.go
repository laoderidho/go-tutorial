package main

// Scan berfungsi Untuk mengambil inputan dari user

import "fmt"

func main() {
	var name string;

	fmt.Print("Masukkan Namamu: ")
	fmt.Scan(&name)
	if name == "Ridho"{
		fmt.Println("Hello", name, "good job")
	}else{
		fmt.Println("your not my bos")
	}
}