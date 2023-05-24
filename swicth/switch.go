package main

import "fmt"

func main() {
	name := "trian"
	switch name {
	case "trian":
		fmt.Println("hallo, trian")
	case "joko":
		fmt.Println("Hallo, joko")
	case "abdul":
		fmt.Println("Hallo, abdul")
	default:
		fmt.Println("harus kenalan dulu nih")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
