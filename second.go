package main

import "fmt"

var y int

func main() {
	fmt.Println(y)
	a, b := 1, "Hardik"
	fmt.Println(a, b)
	var c float32 = 22.34
	fmt.Printf("%v is the value of c of type %T\n", c, c)
	e, f, _ := 12, 13, 14
	fmt.Println(e, f)
}
