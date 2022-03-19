package main
import "fmt"
func main(){
	names := [...]string{"Fadli","udin","Nela"}
	friends := [3]string{"you","and","their"}
	slice := names[:]
	

	fmt.Println(names)
	fmt.Println(friends)
	fmt.Println(slice)
}