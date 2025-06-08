package world

import "testing"

func TestGridSetGet(t *testing.T) {
	g := NewGrid(3, 3)
	if err := g.Set(1, 1, Obstacle); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	state, err := g.Get(1, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if state != Obstacle {
		t.Fatalf("expected Obstacle, got %v", state)
	}
}

func TestGridOutOfBounds(t *testing.T) {
	g := NewGrid(2, 2)
	if err := g.Set(3, 0, Obstacle); err != ErrOutOfBounds {
		t.Fatalf("expected ErrOutOfBounds on set, got %v", err)
	}
	if _, err := g.Get(-1, 0); err != ErrOutOfBounds {
		t.Fatalf("expected ErrOutOfBounds on get, got %v", err)
	}
}

func TestGridInBounds(t *testing.T) {
	g := NewGrid(2, 2)
	if !g.InBounds(1, 1) {
		t.Fatalf("expected (1,1) to be in bounds")
	}
	if g.InBounds(2, 0) {
		t.Fatalf("expected (2,0) to be out of bounds")
	}
}
