package main

import "fmt"


type Blacklist func(string)bool
func main(){
	fmt.Println(registerUser("Reza",func(name string)bool{
		return name != "admin";
	}))
}

func registerUser(name string,blacklist Blacklist)string{
	if blacklist(name){
		return "You are Blocked " + name
	}else{
		return "Well Come "+  name
	}
}