package main

// import required modules
import (
	"fmt"
)

// declare variables and define array content
var strarray = []string{"lorem", "ipsum", "dolor", "sit", "amet"}
var intarray = []int{1, 2, 4, 8, 16}

//var mapone = map[int]string{}
//var maptwo = map[string]interface{}{}

func main() {

	for i := 0; i < len(intarray); i++ {
		fmt.Println(intarray[i])
	}

	for i, v := range strarray {
		fmt.Println("index:", i, "value:", v)

	}

}
