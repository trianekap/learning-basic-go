package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR!!!")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("trian")
}
