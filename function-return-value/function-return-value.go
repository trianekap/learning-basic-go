package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello, bro"
	} else {
		return ("hello " + name)
	}
}

func main() {
	result := getHello("trian")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
