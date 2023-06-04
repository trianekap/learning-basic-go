package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "string")
	case int:
		fmt.Println("Value", value, "int")
	default:
		fmt.Println("Unknown")
	}

}
