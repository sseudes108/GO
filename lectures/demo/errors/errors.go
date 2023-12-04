package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) get(index int) (int, error) {

	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no elements in the index %v", index))
	} else {
		return s.values[index], nil
	}
}

func (s *Stuff) add(value int) {
	s.values = append(s.values, value)
}

func main() {
	newStuff := Stuff{}

	newStuff.add(10)
	newStuff.add(20)
	newStuff.add(30)

	value, err := newStuff.get(5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value is:", value)
	}
}
