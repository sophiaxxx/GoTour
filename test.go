package main

import "fmt"

var x, y, z int
var c, python, java bool

var x1, y1, z1 int = 1, 2, 3
var c1, python1, java1 = true, false, "no!"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(Add(3, 2))
	fmt.Println(x, y, z, c, python, java)
	fmt.Println(x1, y1, z1, c1, python1, java1)
}
