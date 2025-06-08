package main

import (
	"fmt"

	"evogarden/world"
)

// main runs a simple simulation demonstrating agent movement.
func main() {
	g := world.NewGrid(5, 5)
	// place food in the grid
	_ = g.Set(2, 2, world.Food)

	a := &world.Agent{
		X:        0,
		Y:        0,
		Grid:     g,
		Energy:   10,
		Vision:   10,
		MoveCost: 1,
		FoodGain: 5,
	}

	for step := 0; step < 10 && a.Energy > 0; step++ {
		if err := a.MoveTowardsFood(); err != nil {
			fmt.Printf("move failed: %v\n", err)
			break
		}
		fmt.Printf("Step %d:\n%s\nEnergy: %d\n\n", step+1, world.Render(g, []*world.Agent{a}), a.Energy)
	}
}
