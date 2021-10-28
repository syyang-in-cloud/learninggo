package main

import "fmt"

func main() {
	var x int = 10
	var y = 10
	var z int

	var (
		a    int
		b        = 20
		c    int = 30
		d, e     = 40, "hello"
		f, g string
	)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
