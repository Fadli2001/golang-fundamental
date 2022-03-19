package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. "+man.Name	
}

func main(){
	fadli := Man{"Fadli"}

	fadli.Married();
	fmt.Println(fadli)
}

// ketika membuat method struct , maka struct tersebut menajdi pass by value,
// maka ketik merubah filed dari vriable fadli, itu tidka akan berubah jika tidak menggunakan pointer 