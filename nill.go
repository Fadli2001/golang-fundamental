package main

import "fmt"

func newMap(name string)map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"Name" : name,
		}
	}
}

func main(){
	var person map[string]string = nil;
	fmt.Println(person)

	 woman := newMap("Mery")

	if woman == nil{
		fmt.Println("Data Kosong")
	}else{
		fmt.Println(woman)
	}

	// fmt.Println(newMap("Fadli"))
}