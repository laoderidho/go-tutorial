package main

import "fmt"

func main(){

	// ini adalah pemanggilan struct Hello
	h:= Hello{}


	name:= "Ridho Fahreza"

	// ini memanggil method sayHello di struct Hello
	fmt.Println(h.sayHello(name))
}

type Hello struct{

}

// membuat function method di struct
func (h Hello) sayHello(name string)string{
	return "hello " + name
}