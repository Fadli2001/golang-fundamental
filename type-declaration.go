package main;
import "fmt";

func main(){
	type NoKtp string
	type Married bool
	
	var noKtpMe NoKtp = "12312213";
	var status Married = true;

	fmt.Println(noKtpMe);
	fmt.Println(status);


}