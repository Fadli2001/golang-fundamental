package main

import (
	"container/list"
	"fmt"
)

func main(){
	data := list.New()

	data.PushBack("Fadli")
	data.PushFront("Dedi")
	data.PushBack("Rahman")
	data.PushBack("Fauzan")

	// fmt.Println(data.Front().Value)	
	// fmt.Println(data.Front().Next().Value)
	// fmt.Println(data.Front().Next().Value)

for e := data.Front();e != nil ;e = e.Next(){
	fmt.Println(e.Value)
}

}