package main

import "fmt"

type Address struct{
	City,Province,Country string
}

func changeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main(){
 myAddress := Address{"Cianjur","Jawa Barat",""}

 changeCountryToIndonesia(&myAddress)
 fmt.Println(myAddress)
}