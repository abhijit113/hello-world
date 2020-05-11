package main

import "fmt"

func main() {
	fmt.Println("hi....I here....doing control check")
	foo()
	fmt.Println("This will print after foo")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}

func foo() {
	fmt.Println("I am at function foo")
}
