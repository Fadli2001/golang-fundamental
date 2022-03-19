package main

import "fmt"

type Filter func(string)string

func main(){
	
	apologize("Udin","Hasan",forGiveMe);

}

func apologize(name string,object string,filter Filter){
	fmt.Println(name, filter(object));
}

func forGiveMe(name string)string{
	return "maafkan " + name
}