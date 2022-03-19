package main

import "fmt"

type Customer struct{
	name,address string
	age int
	isWorked,isMarried bool
}

func (customer Customer)sayHello (name string){
 fmt.Println("Hello",name ,"My name is ",customer.name)
}

func main(){
	var customer Customer	
	customer.name = "Fadli"
	customer.address = "Cianjur"
	customer.age = 20
	customer.isWorked = false
	customer.isMarried = false	

	customer.sayHello("Udin")
	
	
	// struct literal 
	// Dedi := Customer{
	// 	name : "Dedi",
	// 	address : "Bandung",
	// 	age : 25,
	// 	isWorked : true,
	// 	isMarried: false,
	// }

	// fmt.Println(Dedi)

	// var budi Customer = Customer{"Budi","Surabaya",18,true,false}

	// fmt.Println(budi)
}
