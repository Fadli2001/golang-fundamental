package main

import "fmt"

func main(){
	 months := [...]string{
		 "Januari",
		 "Febuari",
		 "Maret",
		 "April",
		 "Mei",
		 "Juni",
		 "Juli",
		 "Agustus",
		 "September",
		 "Oktober",
		 "November",
		 "Desember",
	 }	 


	 var monthSlice = months[0:2]
	 monthSlice[0] = "Januari Update";	
	 
		
	days := [...]string{
		 "Senin","Selasa","Rabu","Kamis","Jum'at","Sabtu","Minggu",
	 }

	daysSlice := days[2:7]	 
	daysSlice[0] = "Rabu Update" // merubah nama hari rabu	

	
	appendDays := append(daysSlice,"Minggu Lagi")	
	fmt.Println(days); //[Senin, Selasa, Rabu Update, Kamis, Jum'at, Sabtu, Minggu]
	fmt.Println(daysSlice); //[Rabu Update, Kamis, Jum'at, Sabtu, Minggu]
	fmt.Println(appendDays) //[Rabu Update, Kamis, Jum'at, Sabtu, Minggu, Minggu Lagi]

	fmt.Println(cap(daysSlice))
	fmt.Println(cap(appendDays))

	//  membuat slice sekaligus array 

	newSlice := make([]int,2,5)

	newSlice[0] = 1
	newSlice[1] = 2
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice));
	fmt.Println(cap(newSlice));
	
	copySlice := make([]int,2,5)
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	sugar := []string{"teh","jus","nutrisari","air putih","Boba"} //slice


	brandPhone := [...]string {"Samsung","Xiomi","Apple","Fujitsu","Oppo"} //array

	sliceBrandPhone := brandPhone[0:3]; // ["Samsung","Xiomi","Apple"]
	



	fmt.Println(sliceBrandPhone)
	fmt.Println(sugar)
	fmt.Println(brandPhone)


}