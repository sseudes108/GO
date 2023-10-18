package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits++
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)

	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println("Phrase", phrase)

	fmt.Println("first", hello, world)

	replace(&phrase[1], "Go", &counter)

	fmt.Println(phrase)

}
