package main

import "fmt"

func main(){

	// name,sum := sum("Udin",1,2,3,4)

	// fmt.Println(name,sum)


	xSlice := []int{1,2,3,4,5}

	_,sumSlice := sum("Budi",xSlice...)

	fmt.Println(sumSlice)


}

func sum(name string, numbers ...int)(string,int){
	result :=0;
	for x :=0; x <len(numbers); x++{
		result += numbers[x] 
	}
	return name,result
}