package main

import "fmt"

func main(){

	firstName,age := person("Budi",18)

	fmt.Println(firstName,age)
	
}

func person(firstName string,age int)(string,int){
	return firstName,age
}