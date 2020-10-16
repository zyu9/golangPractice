package main

import (
	"fmt"
)

const (
	//a = iota
	//b
	//c
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

//const (
//	a2 = iota
//)

func main() {
	//const myConst int = 43
	//fmt.Printf("%v, %T\n", myConst, myConst)

	//fmt.Printf("%v, %T\n", a, a) /0, int

	//fmt.Printf("%v\n", a) /0
	//fmt.Printf("%v\n", b) /1
	//fmt.Printf("%v\n", c) /2

	//fmt.Printf("%v\n", a2) /0
	fmt.Printf("%v\n", catSpecialist) //6
}
