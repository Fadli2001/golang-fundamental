package main

import (
	"belajar-golang-dasar/database"
	"fmt"
	_ "belajar-golang-dasar/helper"
)


func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}