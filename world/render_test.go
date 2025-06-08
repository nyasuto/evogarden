package world

import "testing"

func TestRender(t *testing.T) {
	g := NewGrid(2, 2)
	_ = g.Set(1, 0, Obstacle)
	_ = g.Set(0, 1, Food)
	a := &Agent{X: 0, Y: 0, Grid: g}

	expected := "A#\nF."
	result := Render(g, []*Agent{a})
	if result != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, result)
	}
}
