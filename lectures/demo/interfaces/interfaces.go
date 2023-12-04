package main

import "fmt"

/*type Reseter interface {
	reset()
}

type InfoPrinter interface {
	printInfo()
}

type Player struct {
	name      string
	health    int
	positionX int
	positionY int
}

type Enemy struct {
	name   string
	health int
	mana   int
}

func (e *Enemy) reset() {
	e.health = 1080
	e.mana = 324
}

func (e *Enemy) printInfo() {
	fmt.Println("Enemy")
	fmt.Println(e.name)
	fmt.Println("Health:", e.health)
	fmt.Println("Mana:", e.mana)
}

func (p *Player) reset() {
	p.health = 100
	p.positionX = 0
	p.positionY = 0
}

func (p *Player) printInfo() {
	fmt.Println(p.name)
	fmt.Println("Health:", p.health)
	fmt.Println("Position X:", p.positionX, "Position Y:", p.positionY)
}

func reset(r Reseter) {
	r.reset()
}
func printInfo(i InfoPrinter) {
	i.printInfo()
}

func resetWithPenalty(r Reseter) {
	if player, ok := r.(*Player); ok {
		player.health = 50
	} else {
		r.reset()
	}
}

func main() {
	player := Player{"Sheila", 50, 5, 5}
	printInfo(&player) // 50,5,5
	resetWithPenalty(&player)
	printInfo(&player) // 100,0,0

	enemy := Enemy{"Danny", 540, 216}
	printInfo(&enemy) // 540,216
	resetWithPenalty(&enemy)
	printInfo(&enemy) // 1080,324
}*/

type Preparer interface {
	prepareDish()
}

type Chicken string
type Salad string

func (c Chicken) prepareDish() {
	fmt.Println("Cook chicken")
}

func (s Salad) prepareDish() {
	fmt.Println("Chop salad. Add dressing")
}

func prepareMultipleDishes(dishes []Preparer) {
	fmt.Println("Preparing dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Println("--Dish:", dish, "--")
		dish.prepareDish()
	}
}

func main() {
	dishes := []Preparer{Chicken("Chiken Wings"), Salad("Iceberg Salad"), Salad("Ceasar's Salad")}
	prepareMultipleDishes(dishes)
}
