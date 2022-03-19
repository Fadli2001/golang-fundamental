package main

import "fmt"


func main(){

	runApp(0)

}

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApp(num int){
	defer logging()
	fmt.Println("Run Application")
	result := 10 /num
	fmt.Println(result)
	
}