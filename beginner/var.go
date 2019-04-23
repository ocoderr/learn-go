package main

// import required modules
import (
	"fmt"
)

// declare variables
var variable string

func main() {

	// define content of variable
	variable = "Lorem Ipsum Dolor Sit Amet"
	x := 123 // 注意检查,是定义新局部变量,还是修改全局变量。该方方式容易造成错误。
	fmt.Println(x)
	// print variable
	fmt.Println(variable)
}
