package main

import (
	"fmt"
)

func main() {
	//var n bool = true
	//fmt.Printf("%v, %T\n", n, n)

	//uint16
	//complex64 (3+7.2i) real(n) imag(n)

	//n := 1 == 1
	//m := 1 == 2
	//fmt.Printf("%v, %T\n", n, n)
	//fmt.Printf("%v, %T\n", m, m)

	a := 10             //1010
	b := 3              //0011
	fmt.Println(a & b)  //0010
	fmt.Println(a | b)  //1011
	fmt.Println(a ^ b)  //1001
	fmt.Println(a &^ b) //0100

	c := 8              //2^3
	fmt.Println(c << 3) //2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) //2^3 / 2^3 = 2^0 = 1

	s := "this is a string"
	t := []byte(s)
	fmt.Printf("%v, %T\n", t, t)

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) //97,  int32
}
