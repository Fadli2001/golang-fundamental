package main

import "fmt"

func main(){
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(5))
}

func factorialLoop(num int)int{
	result := 1;
	for i := num; i>0;i--{
		result *= i;
	}

	return result;
} 

func factorialRecursive(num int)int{
	if num ==1{
		return 1;
	}else{
		return num*factorialRecursive(num-1)
	}
}