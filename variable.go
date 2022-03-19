package main

import "fmt"

func main(){
	var country string
	country = "Indonesia"
	
	city := "Jakarta"

	fmt.Println(country)
	fmt.Println(city)	

	// multiple variable

	var(
		fullName = "Full Name"
		address = "Jakarta"
		age = 12
	)

	fmt.Println(fullName ," " , address , " " ,age)
}