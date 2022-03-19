package main

import "fmt"


func main(){
	runApp(false)
	fmt.Println("Hai")
}

func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("Error dengan message : " , message) 
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(err bool){

	defer endApp() 
	if err {
		panic("Aplikasi error")
	}
	
	
	fmt.Println("Aplikasi berjalan")
}