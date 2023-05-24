package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke ", counter)
	}

	slice := []string{"trian", "eka", "putra", "joko", "karissa"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index ke", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "trian"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
