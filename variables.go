package main

import (
	"fmt"
)

func main() {

	var a = "initial"
	fmt.Println(a)

	var b,c int = 1, 2		// can also be written as b,c := 1,2
	fmt.Println(b,c)

	var d = true
	var e int			//Variables declared without a corresponding initialization are zero-valued
	fmt.Println(d, e)

	f := "short"
	fmt.Println(f)
}
