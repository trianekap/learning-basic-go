package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
