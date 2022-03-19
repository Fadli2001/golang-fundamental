package main

import "fmt"

func main(){

	slice := []int{1,2,3,4}
	x := getSum(slice...);	
	fmt.Println(x)

}

func getSum(numbers ...int)int{
	result := 0;
	for _,n :=range numbers{
		result += n;
	}

	return result;
}