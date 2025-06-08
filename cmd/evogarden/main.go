package main

import (
	"fmt"

	"evogarden/world"
)

// main runs a simple simulation demonstrating agent movement.
func main() {
	g := world.NewGrid(20, 20)
	// place food in the grid
	_ = g.Set(10, 10, world.Food)

	agents := []*world.Agent{
		{
			X:        0,
			Y:        0,
			Grid:     g,
			Energy:   20,
			Vision:   30,
			MoveCost: 1,
			FoodGain: 5,
		},
		{
			X:        19,
			Y:        0,
			Grid:     g,
			Energy:   20,
			Vision:   30,
			MoveCost: 1,
			FoodGain: 5,
		},
		{
			X:        0,
			Y:        19,
			Grid:     g,
			Energy:   20,
			Vision:   30,
			MoveCost: 1,
			FoodGain: 5,
		},
		{
			X:        19,
			Y:        19,
			Grid:     g,
			Energy:   20,
			Vision:   30,
			MoveCost: 1,
			FoodGain: 5,
		},
		{
			X:        10,
			Y:        0,
			Grid:     g,
			Energy:   20,
			Vision:   30,
			MoveCost: 1,
			FoodGain: 5,
		},
	}
	sim := &world.Simulation{Grid: g, Agents: agents}

	for step := 0; step < 20; step++ {
		sim.Step()
		fmt.Printf("Step %d:\n%s\n\n", step+1, world.Render(g, sim.Agents))
	}
}
