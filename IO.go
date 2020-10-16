package main

// Importing packages
import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	// Reading a file (input operation)
	Character_list, err := ioutil.ReadFile("input.txt")

	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
		return
	}
	// for loop to count the number of characters using counter
	var counter int
	for range Character_list {
		counter++
	}

	// Printing the number of characters
	fmt.Println(" Number of characters in a string:", counter)

	// Converting counter of type int into string first and then into byte array
	count := []byte(strconv.Itoa(counter))

	// Writing a file (output operation)
	err1 := ioutil.WriteFile("output.txt", count, 0777)

	// handle this error
	if err1 != nil {
		// print it out
		fmt.Println(err1)
		return
	}
}
