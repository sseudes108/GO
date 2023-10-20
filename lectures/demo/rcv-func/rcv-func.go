package main

import "fmt"

type Space struct {
	occupied bool
}
type ParkinLot struct {
	spaces []Space
}

func occupySpace(lot *ParkinLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkinLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkinLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkinLot{spaces: make([]Space, 5)}
	fmt.Println(lot)

	lot.occupySpace(1)
	occupySpace(&lot, 3)

	fmt.Println(lot)
	lot.vacateSpace(3)

	fmt.Println(lot)
}
