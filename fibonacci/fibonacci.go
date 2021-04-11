package main

import "fmt"

// a function that returns the next fibonacci number
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		temp := f1 + f2
		temp2 := f1
		f1 = f2
		f2 = temp
		return temp2
	}
}

func main() {
	fmt.Println("First 10 fibonacci numbers ### ")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
