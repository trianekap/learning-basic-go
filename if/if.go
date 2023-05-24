package main

import "fmt"

func main() {
	name := "trian"

	if name == "trian" {
		fmt.Println("Hallo,", name)
	} else if name == "joko" {
		fmt.Println("Hallo,", name)
	} else if name == "putra" {
		fmt.Println("Hallo,", name)
	} else {
		fmt.Println("boleh kenalan dulu?")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
