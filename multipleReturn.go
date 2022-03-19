package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main(){
	name,_ := getIdent();

	fmt.Println("Hallo Nama saya ",name)
}

func getIdent()(string,int){
	return "Fadli",20;
}

func bmi()int{

	var height int
	var weight int

	fmt.Println("Enter Your Height :")
	fmt.Scan(&height)
	fmt.Println("Enter Your Weight :")
	fmt.Scan(&weight)
	weightN := weight*weight
	result := height/weightN
	return result;


	
}

