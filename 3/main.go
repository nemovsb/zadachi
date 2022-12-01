package main

import "fmt"

func main() {

	num := 10
	a := func() int {
		//fmt.Printf("a = %d\n", num-1)
		num--
		return num
	}

	defer func() {
		fmt.Printf("a0 = %d\n", a())
	}()
	defer func() {
		//defer panic(3)
		defer fmt.Printf("a1 = %d\n", a())
		defer fmt.Printf("a2 = %d\n", a())
		defer fmt.Printf("a3 = %d\n", a())
		defer fmt.Printf("a4 = %d\n", a())
		fmt.Printf("a5 = %d\n", a())
	}()
	defer fmt.Printf("a6 = %d\n", a())
	fmt.Printf("a7 = %d\n", a())
}

func count(a int) int {
	//fmt.Printf("a = %d\n", a-1)
	return a - 1
}
