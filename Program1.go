package main

import (
	"fmt"
)

type profile struct {
	name         string
	company_name string
	age          int
}

func add(v profile) {

	fmt.Println("name:", v.name)
	fmt.Println("Company_name", v.company_name)
	fmt.Println("age:", v.age)
}

//Ceate an "Add" function for adding the details and finally print it.

func main() {
	var profile1 profile
	profile1.name = "Nandini"
	profile1.company_name = "Medivo"
	profile1.age = 30
	add(profile1)

}
