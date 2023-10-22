package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (chicken Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}
func (salad Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf(("--Dish: %v--\n"), dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("chicken wings"), Salad("iceberg salad")}
	prepareDishes(dishes)
}
