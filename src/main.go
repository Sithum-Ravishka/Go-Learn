package main

import (
	"fmt"
)

func main() {
	var std1 = make(map[string]string)

	std1["name"] = "Sithum"
	std1["age"] = "25"

	std2 := make(map[string]int)
	std2["grade"] = 11
	std2["age"] = 22

	fmt.Printf("%v\n", std1)
	fmt.Printf("%v\n", std2)

	fmt.Printf(std1["name"])

	std2["grade"] = 90 // Update element in map
	std2["class"] = 10 // Add new element in map
	fmt.Printf("\nUpdated Map: %v\n", std2)

	delete(std2, "grade") // Remove element in map
	fmt.Printf("\nUpdated Map After Delete: %v\n", std2)

	val1, ok1 := std1["brand"] //Check For Specific Elements in a Map
	val2, ok2 := std2["class"]
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)

}
