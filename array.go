package main

import "fmt"

func main(){

	// deklarasi array disertakan jumlah maksimum data dan tipe data 
	var days [7]string;

	days[0] = "Senin"
	days[1] = "Selasa"
	days[2] = "Rabu"

	// Gaya Horizontal 
	var months = [12]string{"Januari","Febuari","Maret"}

	// Gaya Vertikal
	var food = [4]string {
		"Sop Ayam",
		"Ayam Goreng",
		"Nasgor",
		"Mie",
	}

	//tanpa menuliskan jumlah data, namun harus sekaligus mengisi nilainya
	var number = [...]int{1,2,3,4,5}


	fmt.Println(months)
	fmt.Println(number)
	fmt.Println(food)
	


} 