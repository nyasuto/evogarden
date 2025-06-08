package world

import "testing"

func TestSimulationStep(t *testing.T) {
	g := NewGrid(3, 3)
	_ = g.Set(2, 2, Food)
	a1 := &Agent{X: 0, Y: 0, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3, Vision: 10}
	a2 := &Agent{X: 1, Y: 0, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3, Vision: 10}
	sim := &Simulation{Grid: g, Agents: []*Agent{a1, a2}}

	sim.Step()

	if a1.Energy >= 5 {
		t.Fatalf("expected a1 energy to decrease")
	}
	if a2.Energy >= 5 {
		t.Fatalf("expected a2 energy to decrease")
	}
}
