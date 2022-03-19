package main

import "fmt"


func main(){
	runApp(false)
}

func endApp(){
	fmt.Println("Aplikasi selesai")
}

func runApp(err bool){

	defer endApp()
	if err {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi berjalan")
}