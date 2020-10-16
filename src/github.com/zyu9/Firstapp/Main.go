package main

import (
	"fmt"
	//"strconv"  | string convertion
	//in method: j = strconv.Itoa(i)
)

var i int = 42

//var{
//	actorName string = "Elisabeth Sladen"
//	companion string = "Sarah Jane Smith"
//	doctorName int = 3
//	season int = 11
//}

//var{
//	counter int = 0
//}

func main() {
	//	var i int = 42
	// var j float32
	// j = float32(i)

	// k := 99.
	//fmt.Printf("%v, %T\n", i, i)
	fmt.Println(i)
	var i int = 27 //shadow
	fmt.Println(i)
}
