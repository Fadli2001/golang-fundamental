package main

import "fmt"

func main(){

	goodbye := getGoodby;	
	fmt.Println(goodbye("Udin"))	
	value1 := plus(2)
	result := value1 + 3
	fmt.Println(result)
}

func getGoodby(name string)string{
	return "GoodBye "+name
}

func plus(a int)int{
 return a +1 
}
