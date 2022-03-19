package main

import "fmt"

type Customer struct{
	name,address string
	age int
	isWorked,isMarried bool
}

func main(){

	var customer Customer	
	customer.name = "Hasan"
	customer.address = "Surabaya"
	customer.age = 20
	customer.isWorked = false
	customer.isMarried = false	

	
	// struct literal 
	Dedi := Customer{
		name : "Dedi",
		address : "Bandung",
		age : 25,
		isWorked : true,
		isMarried: false,
	}

	fmt.Println(Dedi)

	var budi Customer = Customer{"Budi","Surabaya",18,true,false}

	fmt.Println(budi)
}
