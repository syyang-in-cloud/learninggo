package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)

}
