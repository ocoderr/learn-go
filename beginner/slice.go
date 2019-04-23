package main

// import required modules
import (
	"fmt"
)

// main function
func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Deleting from a slice is basically appending the entire slice to itself, but excluding the deleted item
	mySlice = append(mySlice[:4], mySlice[5:]...)
	fmt.Println(mySlice)

	var str string = "Lorem ipsum dolor sit amet"
	fmt.Println(str[6:11])
}
