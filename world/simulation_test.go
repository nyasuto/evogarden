package world

import "testing"

func TestSimulationStep(t *testing.T) {
	g := NewGrid(20, 20)
	_ = g.Set(10, 10, Food)
	agents := []*Agent{
		{X: 0, Y: 0, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3, Vision: 20},
		{X: 1, Y: 0, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3, Vision: 20},
		{X: 0, Y: 1, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3, Vision: 20},
		{X: 19, Y: 19, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3, Vision: 20},
		{X: 10, Y: 0, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3, Vision: 20},
	}
	sim := &Simulation{Grid: g, Agents: agents}

	sim.Step()

	for i, a := range agents {
		if a.Energy >= 5 {
			t.Fatalf("expected agent %d energy to decrease", i)
		}
	}
}
