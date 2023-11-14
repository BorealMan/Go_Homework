package main

/*
	Written By RK

	Task:
		- Simulate The Moose Crossing The Bridge Without Falling To Their Death!
		- Only Allow A Moose To Start Crossing Every 1 Second
		- You Must Use A gorountine

*/

import (
	hw "app/homework"
	"fmt"
	"time"
)

func main() {

	// Create Bridge
	Bridge := hw.CreateBridge()

	// Generate Mooses With Random Speed & Weight
	Mooses, err := hw.GenerateMooses(50)

	if err != nil {
		panic(err)
	}
	hw.PrintMooses(Mooses)

	index := len(Mooses) - 1
	for {
		if Mooses == nil && len(Bridge.Load) == 0 {
			fmt.Println("All The Mooses Successfully Crossed!")
			break
		}
		/*
			Load The Bridge Up To Capacity
				Tip:
					- Use The Built In Calc To Check The Load On The Bridge And The Wieght of the Next Moose
					- If That If Less Than The Bridge Capacity Then Use MooseCross
		*/
		for {
			if index == -1 {
				Mooses = nil
				break
			}
			// Put Your Code Here
		}
		// Simulate Bridge
		err := Bridge.Simulate()
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
