package world

import "bytes"

// Render returns an ASCII representation of the grid with the provided agents overlaid.
// Empty cells are '.', obstacles '#', food 'F' and agents 'A'.
func Render(g *Grid, agents []*Agent) string {
	agentMap := make(map[int]bool)
	for _, a := range agents {
		if g.InBounds(a.X, a.Y) {
			agentMap[g.index(a.X, a.Y)] = true
		}
	}
	var buf bytes.Buffer
	for y := 0; y < g.Height; y++ {
		if y > 0 {
			buf.WriteByte('\n')
		}
		for x := 0; x < g.Width; x++ {
			idx := g.index(x, y)
			if agentMap[idx] {
				buf.WriteByte('A')
				continue
			}
			switch g.cells[idx] {
			case Empty:
				buf.WriteByte('.')
			case Obstacle:
				buf.WriteByte('#')
			case Food:
				buf.WriteByte('F')
			}
		}
	}
	return buf.String()
}
