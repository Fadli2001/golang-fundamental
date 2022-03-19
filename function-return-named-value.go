package main

import "fmt"

func main(){

	a,b,c := getFullName();
	fmt.Println(a,b,c)
}


func getFullName()(firstName,middleName,lastName string){
	fmt.Println("Enter Your First Name :")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your Middle Name :")
	fmt.Scan(&middleName)
	fmt.Println("Enter Your Last Name :")
	fmt.Scan(&lastName)

	return

}