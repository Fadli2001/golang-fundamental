package main

import "fmt"

func random() interface{}{
	return "sdf"
}

func main(){

	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type){
	case string :
		fmt.Println("Value", value ,"is String")
	case int:
		fmt.Println("Value", value ," is Integer")
	default:
		fmt.Println("unknown")
	}
 
}