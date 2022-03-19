package main

import "fmt"

func main(){

	show := getData("Fadli");

	fmt.Println(show)



}

func getData(name string)string{
	return "Hallo nama saya "+name
}

func sayUra(){
	fmt.Println("Uraaa")
}