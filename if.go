package main

import "fmt"

func main(){

	name := "Udin"

	if name == "Udin"{
		fmt.Println("Hallo " + name)
	}else if name == "Upin"{
		fmt.Println("Hallo " + name)
	}else{
		fmt.Println("Siapa ya ?")
	}

	// if short statement 

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	}else{
		fmt.Println("It's oke")
	} 


}