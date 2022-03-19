package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		if i == 5{
			break
		}
		fmt.Println("Perulangan ke-",i+1)
	}

	for i := 0; i<10;i++{

		if i == 2{
			continue
		}
		fmt.Println("Perulangan Continue ke-",i)
	}
}