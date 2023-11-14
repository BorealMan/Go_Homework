package homework

/*
	Written By RK

	Task:
		- Create Random Mooses With A Random Weight Between 800 to 1,600 and Random Speed Between 2-10
		- Run Simulation of The Mooses Crossing The Bridge Without Falling To Their Death!

	Tips:
		- Only Allow A Moose To Start Crossing Every 1 Second
*/

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Moose struct {
	Id      int
	Weight  int
	Speed   int
	Crossed bool
}

type Bridge struct {
	Capacity int
	Load     []*Moose
}

func CreateMoose(weight int, speed int) (*Moose, error) {
	if weight <= 0 {
		return nil, errors.New("Moose Weight Must Be Greater Than Zero!")
	}
	if speed <= 0 {
		return nil, errors.New("Moose Speed Must Be Greater Than Zero!")
	}
	moose := new(Moose)
	moose.Weight = weight
	moose.Speed = speed
	moose.Crossed = false
	return moose, nil
}

/*
Sourced From https://gosamples.dev/random-numbers/
*/
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func GenerateMooses(n int) ([]*Moose, error) {
	var Mooses []*Moose
	for i := 0; i < n; i++ {
		// Create Moose
		moose, err := CreateMoose(RandInt(800, 1600), RandInt(3, 10))
		if err != nil {
			panic(err)
		}
		// Add To Mooses
		Mooses = append(Mooses, moose)
	}
	return Mooses, nil
}

func PrintMooses(Mooses []*Moose) {
	for _, moose := range Mooses {
		fmt.Print(*moose, " ")
	}
	fmt.Println()
}

func CreateBridge() *Bridge {
	bridge := new(Bridge)
	bridge.Capacity = 10000
	return bridge
}

// A Moose Attempts To Cross The Bridge
func (b *Bridge) MooseCross(moose *Moose) {
	b.Load = append(b.Load, moose)
	fmt.Printf("A %d Moose Has Started Crossing - Total Weight: %d\n", moose.Weight, b.CalcLoad())
	for {
		fmt.Printf("A %d Moose Has Moved!\n", moose.Weight)
		if moose.Speed <= 0 {
			moose.Crossed = true
			break
		}
		moose.Speed -= 1
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("A %d Moose Has Finished Crossing - Total Weight: %d\n", moose.Weight, b.CalcLoad())
}

// Calculate Load
func (b *Bridge) CalcLoad() int {
	load := 0
	for _, moose := range b.Load {
		if !moose.Crossed {
			load += moose.Weight
		}
	}
	return load
}

// Calculates If Load Weight is Greater Than Capacity
func (b *Bridge) IsBridgeSturdy() bool {
	load := b.CalcLoad()
	if load > b.Capacity {
		return false
	}
	return true
}

func (b *Bridge) RemoveMooseFromLoad() {
	for i, moose := range b.Load {
		// Moose Finished Crossing, Remove From Load
		if moose.Crossed {
			if i == len(b.Load)-1 {
				b.Load = b.Load[:i]
			} else {
				b.Load = append(b.Load[:i], b.Load[i+1:]...)
			}
		}
	}
}

// Runs Moose Simulation
func (b *Bridge) Simulate() error {
	// End Simulation If Empty Load
	if len(b.Load) <= 0 {
		return nil
	}
	isSturdy := b.IsBridgeSturdy()
	if !isSturdy {
		PrintMooses(b.Load)
		return fmt.Errorf("The Bridge Has Broke! %d Moose Fell To Their Death!\nThe Load Was: %d\n", len(b.Load), b.CalcLoad())
	}
	b.RemoveMooseFromLoad()
	return nil
}
