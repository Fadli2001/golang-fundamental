package main

import "fmt"



func main(){

	genap  := 2	

	 var pointerGenap *int = &genap	

	fmt.Println(pointerGenap) //0xc0000b8000
	fmt.Println(*pointerGenap) //2

}