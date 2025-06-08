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

	agents := []*world.Agent{
		{
			X:        0,
			Y:        0,
			Grid:     g,
			Energy:   10,
			Vision:   10,
			MoveCost: 1,
			FoodGain: 5,
		},
		{
			X:        4,
			Y:        4,
			Grid:     g,
			Energy:   10,
			Vision:   10,
			MoveCost: 1,
			FoodGain: 5,
		},
	}
	sim := &world.Simulation{Grid: g, Agents: agents}

	for step := 0; step < 10; step++ {
		sim.Step()
		fmt.Printf("Step %d:\n%s\n\n", step+1, world.Render(g, sim.Agents))
	}
}
