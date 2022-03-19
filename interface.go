package main

import "fmt"


// type Person struct{
// 	Name string
// }

// func (person Person) GetName()string{
//  return person.Name
// }

type HashName interface{
	GetName() string
}

type Car struct{
	Name string
	cc int
}

func (car Car) GetName()string{
	return car.Name
}

func sayHello(hasName HashName){
	fmt.Println("Hello",hasName.GetName())
}

func main(){

	toyota := Car{"Tyota",12}
	sayHello(toyota)

	

	// fadli := Person{
	// 	Name : "Fadli",
	// }

	// var udin Person
	// udin.Name = "Udin"


	/** mengapa var fadli bisa di jadikan value parameter sayHello yang memiliki type HashName interface, 
	karena struct menurpakan implemen dari interface,yang dimana struct itu sendiri memiliki sebuah func yang sama yaitu GetName()*/

	// sayHello(fadli)
	// sayHello(udin)

}