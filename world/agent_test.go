package world

import "testing"

func TestAgentSearchFood(t *testing.T) {
	g := NewGrid(3, 3)
	if err := g.Set(2, 2, Food); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	a := &Agent{X: 0, Y: 0, Grid: g, Vision: 10}
	x, y, found := a.SearchFood()
	if !found {
		t.Fatalf("expected to find food")
	}
	if x != 2 || y != 2 {
		t.Fatalf("expected to find food at (2,2), got (%d,%d)", x, y)
	}
}

func TestAgentMoveConsumesEnergy(t *testing.T) {
	g := NewGrid(3, 3)
	a := &Agent{X: 0, Y: 0, Grid: g, Energy: 5, MoveCost: 1}
	if err := a.Move(1, 0); err != nil {
		t.Fatalf("unexpected move error: %v", err)
	}
	if a.Energy != 4 {
		t.Fatalf("expected energy 4, got %d", a.Energy)
	}
}

func TestAgentMoveToFood(t *testing.T) {
	g := NewGrid(3, 3)
	if err := g.Set(1, 0, Food); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	a := &Agent{X: 0, Y: 0, Grid: g, Energy: 5, MoveCost: 1, FoodGain: 3}
	if err := a.Move(1, 0); err != nil {
		t.Fatalf("unexpected move error: %v", err)
	}
	if a.Energy != 7 {
		t.Fatalf("expected energy 7, got %d", a.Energy)
	}
	state, _ := g.Get(1, 0)
	if state != Empty {
		t.Fatalf("expected food to be consumed")
	}
}
