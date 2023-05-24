package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "trian"
	names[1] = "eka"
	names[2] = "putra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		87,
		56,
		45,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println("jumlah array =", len(names))
	fmt.Println("jumlah array =", len(values))
}
