package main

import "fmt"

func main() {
	name := "trian"
	counter := 0

	increment := func() {
		name := "jafar"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)

}
