package main

import "fmt"

func main() {
	var x [3]int
	var y = [3]int{10, 20, 30}
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	var a = [...]int{1, 2, 3}
	var b [2][3]int

	b[0][0] = 1
	b[1][2] = 6

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(len(b[0]))
	fmt.Println(len(&b))
	fmt.Println(len(&b[0]))
}
