package main

import "fmt"

func main() {
	var map1 map[int]string
	map1 = make(map[int]string)
	map1[1] = "Nandini"
	map1[6] = "Sharad"
	map1[4] = "Milin"
	map1[10] = "Pavan"
	map1[2] = "shwetha"
	map1[7] = "raja"

	for id := range map1 {
		fmt.Println("id", id, "value", map1[id])
	}
	fmt.Println("AFTER DELETION Of 3 records")
	delete(map1, 1)
	delete(map1, 4)
	delete(map1, 7)
	for id := range map1 {
		fmt.Println("id", id, "value", map1[id])
	}
	map1[10] = "Maggi"
	fmt.Println("AFTER UPDATION")
	for id := range map1 {
		fmt.Println("id", id, "value", map1[id])
	}
}

