package main

import "fmt"

type Address struct {
	City,Province,Country string
}

func main(){
	var address1 Address = Address{"Jakarta","Jawa Barat","Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Depok"

	//vallue dari address1 tidak ikut berubah
	// address2 = &Address{"cianjur","Jawa Barat","Indonesia"}

	//jika ingin ikut berubah bisa mneggunakan tanda bintang
	*address2 = Address{"cianjur","Jawa Barat","Indonesia"}
	

	fmt.Println(address1)
	fmt.Println(address2)	
	fmt.Println(address3)	

	/**
	melakukan assigmanet address1 menjadi address2 , di golang itu pass by value, dia melakukan duplicate nilainya, 
	bisa di lihat ketika merubah city dari address2 ,city pada address1 tidak ikut berubah

	untuk menanggani hal tersebut bisa di lakukan dengan menggunakan pointer &... > pass by reference

	contoh address2 := &address1
	*/

	address4 := new(Address)
	fmt.Println(address4)
}