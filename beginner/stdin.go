package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter digits to sum separated by a space:")
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err == nil {
		fmt.Println("Here is the result:", sumNumbers(str))
	}

	//Stdin read buffer
	fmt.Println("Enter a string")
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println("you entered \"" + string(line) + "\"")
		for _, char := range string(line) {
			fmt.Println(char)
		}
	}
	/*
		//Scanf to wait for an input
		fmt.Println("Enter a number")
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println("not a number")
		} else {
			fmt.Print(i)
			fmt.Println(" is a number")
		}
	*/
	var s string
	fmt.Print("please insert a string an press enter ")
	fmt.Scan(&s)
	fmt.Printf("read string \"%v\" from stdin\n", s)

}

func sumNumbers(str string) float64 {
	var sum float64
	for _, v := range strings.Fields(str) {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			sum += i
		}
	}
	return sum
}
