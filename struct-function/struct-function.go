package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuu from", a.Name)
}

func main() {
	var Trian Customer
	Trian.Name = "Trian Eka Putra"
	Trian.Address = "Indonesia"
	Trian.age = 28

	Trian.sayHi("joko")
	Trian.sayHuuu()

	// fmt.Println(Trian.Name)
	// fmt.Println(Trian.Address)
	// fmt.Println(Trian.age)

	// joko := Customer{
	// 	Name:    "Joko Tingkir",
	// 	Address: "Cirebon",
	// 	age:     35,
	// }

	// fmt.Println(joko)

	// budi := Customer{"Budi", "Jakarta", 35}
	// fmt.Println(budi)
}
