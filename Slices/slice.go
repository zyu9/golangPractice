package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	d := c[:]   //slice of all elements
	e := c[3:]  //slice from 4th element to end
	f := c[:6]  //slice first 6 elements
	g := c[3:6] //slice the 4th, 5th, and 6th elements
	c[5] = 42
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	h := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))

	i := []int{}
	i = append(i, 1)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	//i = append(i, 2, 3, 4, 5)
	i = append(i, []int{2, 3, 4, 5}...)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))

	j := []int{1, 2, 3, 4, 5}
	//k := j[1:]
	//k := j[:len(j)-1]
	k := append(j[:2], j[3:]...)
	fmt.Println(k)

}
