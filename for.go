package main

import "fmt"

func main(){
	

	// for a := 1; a <= 5; a++{
	// 	fmt.Println("Perulangan ke - ", + a )
	// }

	slice := []string {"Fadli","Rahman","Fauzan","Udin","Sapri"}

	for a := 0;a < len(slice); a++{
		fmt.Println( "slice ke-", a+1 ,"=",slice[a]);
	}

	fmt.Println("================")

	for _, value := range slice{
		fmt.Println( value)
	}
}