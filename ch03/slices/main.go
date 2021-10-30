package main

import "fmt"

func main() {
	var x = []int{10, 20, 30}
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	var z []int

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(x == nil)
	fmt.Println(z == nil)

	x = append(x, 40)
	fmt.Println(x)
	x = append(x, y...)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 50)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	a := make([]int, 5)
	a = append(a, 10)
	fmt.Println(a)
	b := make([]int, 5, 10)
	b = append(b, 20)
	fmt.Println(b)

	var c = []int{}
	fmt.Println(c)

}
