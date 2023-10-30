package main

import "fmt"

const (
	small = iota
	medium
	large
)
const (
	ground = iota
	air
)

type Beltsize int
type Shipping int

func (b Beltsize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Convoyer interface {
	Convoy() Beltsize
}
type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Convoyer
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam Mail"
}

func (s *SpamMail) Ship() Shipping {
	return air
}

func (s *SpamMail) Convoy() Beltsize {
	return small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v\n", item, item.Convoy())
	fmt.Printf("Ship by %v\n", item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return ground
}

func main() {

	sm := SpamMail{40000}
	automate(&sm)
}
