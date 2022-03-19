package main;

import "fmt"

func main(){

	food := "makanan";

	switch food {
	case "makanan":
		fmt.Println("nyamnyam")
	case "minuman":
		fmt.Println("sruput")
	default:
		fmt.Println("not valid")
	}


	// short statement 

	switch positive := age; positive > 0{
	case true :
		fmt.Println("Oke")
	case false:
		fmt.Println("invalid")
	}

length := 2;

	switch{
	case length > 10 && length < 15:
		fmt.Println("Lebih dari 10")
	case length < 10 && length > 5:
		fmt.Println("Kurang dari 10")
	default:
		fmt.Println("Nilai tidak terdaftar")
	}

}