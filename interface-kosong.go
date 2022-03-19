package main

import "fmt"

func ups(i int) (interface{}){
	if i == 1{
		return 1
	}else if i == 2{
		return true
	}else{
		return "ups"
	}
}

func randomParam(value interface{}) interface{}{
	return value
}

func main(){

	var emptyInterface interface{} = ups(1)
	boolInterface := ups(4)

	

	fmt.Println(emptyInterface)
	fmt.Println(boolInterface)
	fmt.Println(randomParam(false))
}