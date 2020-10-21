package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 30
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		break
	case j <= 20:
		fmt.Println("less than or equal to twenty")
		break
	default:
		fmt.Println("greater than twenty")
		break
	} //break fallthrough

	var k interface{} = 1
	switch k.(type) {
	case int:
		fmt.Println("k is an int")
	case float64:
		fmt.Println("k is a float64")
	case string:
		fmt.Println("k is string")
	default:
		fmt.Println("k is another type")
	}
}
