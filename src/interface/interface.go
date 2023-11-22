package main

import "fmt"

type Preparer interface {
	PreparerDish()
}

type Chicken string
type Salad string

func (c Chicken) PreparerDish() {
	fmt.Println("Cook Chicken")
}

func (c Salad) PreparerDish() {
	fmt.Println("Chop salad")
	fmt.Println("Add dressing")
}

func PreparerDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]

		fmt.Printf("---Dish: %v---\n", dish)
		dish.PreparerDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Iceberg Salad")}
	PreparerDishes(dishes)
}
