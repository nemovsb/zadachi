package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(recover()) //1
	}()
	defer func() {
		//defer panic(3)
		defer fmt.Println("1: ", recover())
		defer fmt.Println("2: ", recover())
		defer fmt.Println("3: ", recover()) //2
		defer panic(1)
		fmt.Println("4: ", recover())
		recover()
		fmt.Println("5: ", recover())
	}()
	defer recover()
	panic(2)
}
