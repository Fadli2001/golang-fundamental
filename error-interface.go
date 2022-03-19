package main

import (
	"fmt"
	"errors"
)

func pembagi(nilai int,pembagi int)(int,error){
	if pembagi == 0{
		return 0,errors.New("Pembagi tidak boleh nol")
	}else{
		return nilai/pembagi,nil 
	}
}

func main(){

	salah := errors.New("Error Lagi!!!")

	fmt.Println(salah)
	
	hasil,err := pembagi(10,0)

	if err == nil{
		fmt.Println("hasil",hasil)
	}else{
		fmt.Println(err.Error())
	}


}