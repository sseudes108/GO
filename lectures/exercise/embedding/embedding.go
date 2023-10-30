//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (b *BandwidthUsage) AvarageBandwithUsage() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	fmt.Println("Avarage Bandwith Usage:", sum/len(b.amount))
	return sum / len(b.amount)
}

func (c *CpuTemp) AvarageCpuTemp() float32 {
	var sum float32 = 0.0
	for i := 0; i < len(c.temp); i++ {
		sum += float32(c.temp[i])
	}
	fmt.Println("Avarage Cpu Temperature:", sum/float32(len(c.temp)))
	return sum / float32(len(c.temp))
}

func (m *MemoryUsage) AvarageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	fmt.Println("Avarage Memory Usage:", sum/len(m.amount))
	return sum / len(m.amount)
}

func dashBoardUsageStats(board *Dashboard) {
	board.BandwidthUsage.AvarageBandwithUsage()
	board.CpuTemp.AvarageCpuTemp()

	board.MemoryUsage.AvarageMemoryUsage()
}

func main() {

	board := Dashboard{
		BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}},
		CpuTemp{[]Celcius{50, 51, 53, 51, 52}},
		MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}},
	}

	board.AvarageBandwithUsage()
	board.AvarageCpuTemp()
	board.AvarageMemoryUsage()
	fmt.Println("--------------------------")
	dashBoardUsageStats(&board)
}
