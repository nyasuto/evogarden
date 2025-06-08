package world

import "errors"

// CellState represents the content of a grid cell.
type CellState int

const (
	// Empty indicates no obstacle or entity in the cell.
	Empty CellState = iota
	// Obstacle indicates that the cell is blocked.
	Obstacle
	// Food indicates a food source that agents can consume.
	Food
)

// Grid represents a 2D grid world.
type Grid struct {
	Width  int
	Height int
	cells  []CellState
}

// NewGrid returns a Grid with the given width and height.
func NewGrid(w, h int) *Grid {
	if w <= 0 || h <= 0 {
		panic("grid dimensions must be positive")
	}
	return &Grid{Width: w, Height: h, cells: make([]CellState, w*h)}
}

// index converts 2D coordinates to a slice index.
func (g *Grid) index(x, y int) int {
	return y*g.Width + x
}

// InBounds checks whether coordinates are inside the grid.
func (g *Grid) InBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

// Set updates the state of the cell at (x,y).
func (g *Grid) Set(x, y int, state CellState) error {
	if !g.InBounds(x, y) {
		return ErrOutOfBounds
	}
	g.cells[g.index(x, y)] = state
	return nil
}

// Get returns the state of the cell at (x,y).
func (g *Grid) Get(x, y int) (CellState, error) {
	if !g.InBounds(x, y) {
		return Empty, ErrOutOfBounds
	}
	return g.cells[g.index(x, y)], nil
}

// ErrOutOfBounds is returned when coordinates are outside of the grid.
var ErrOutOfBounds = errors.New("coordinates out of bounds")
