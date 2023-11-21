package main

import "fmt"

func main() {
	var std1 = make(map[string]string)

	std1["name"] = "Sithum"
	std1["age"] = "25"

	std2 := make(map[string]int)
	std2["name"] = 11
	std2["age"] = 22

	fmt.Printf("%v", std1)
	fmt.Printf("%v", std2)
}
