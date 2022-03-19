package main
import "fmt"

func main(){

	nilaiUAS := 50
	nilaiAbsen := 80

	nilaiLulusUAS := nilaiUAS >= 75
	nilaiLulusAbsen := nilaiAbsen >= 75

	var lulus bool = nilaiLulusUAS && nilaiLulusAbsen

	fmt.Println(lulus)
	fmt.Println(nilaiUAS >= 75 && nilaiAbsen >= 74);
	
}