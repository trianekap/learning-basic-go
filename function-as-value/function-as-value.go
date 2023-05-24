package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye

	result := getGoodBye("trian")
	fmt.Println(result)

	fmt.Println(goodbye("trian"))
}
