package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "trian"
	middleName = "eka"
	lastName = "putra"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
