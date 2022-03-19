package main

import "fmt"

func main(){

	person := map[string]string{
		"name" : "Udin",
		"address" : "Jakarta",		
	}

	days := map[int]string{}

	days[1] = "Senin"
	days[2] = "Selasa"
	days[3] = "Rabu"
	days[4] = "Kamis"


	fmt.Println(days)

	delete(person,"address")
	fmt.Println(person)

	school := make(map[string]string)

	school["name"] = "SMKN 1 Cipanas"

	fmt.Println(school)
	nilaiIpa := 76;
	nilaiIps := 75;

	if(nilaiIpa >= 75 && nilaiIps >= 73){
		fmt.Println("lulus")
	}else{
		fmt.Println("Tidak Lulus")
	}

}