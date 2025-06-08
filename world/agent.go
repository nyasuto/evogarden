package world

import "errors"

// Agent represents an entity in the world capable of moving and searching for food.
type Agent struct {
	X, Y     int   // current position
	Energy   int   // remaining energy
	Vision   int   // maximum search distance (0 = unlimited)
	MoveCost int   // energy consumed per move
	FoodGain int   // energy gained when consuming food
	Grid     *Grid // world the agent resides in
}

// ErrBlocked is returned when an agent attempts to move into an obstacle.
var ErrBlocked = errors.New("movement blocked by obstacle")

// Move attempts to move the agent by the given delta. It applies energy costs
// and handles food consumption when the destination contains food.
func (a *Agent) Move(dx, dy int) error {
	nx, ny := a.X+dx, a.Y+dy
	if !a.Grid.InBounds(nx, ny) {
		return ErrOutOfBounds
	}
	state, err := a.Grid.Get(nx, ny)
	if err != nil {
		return err
	}
	if state == Obstacle {
		return ErrBlocked
	}
	a.X = nx
	a.Y = ny
	a.Energy -= a.MoveCost
	if state == Food {
		a.Energy += a.FoodGain
		_ = a.Grid.Set(nx, ny, Empty)
	}
	return nil
}

// SearchFood performs a breadth-first search up to the agent's vision range
// looking for the closest food. It returns the coordinates of the food and a
// boolean indicating whether any food was found.
func (a *Agent) SearchFood() (int, int, bool) {
	type node struct{ x, y, d int }
	visited := make([]bool, a.Grid.Width*a.Grid.Height)
	q := []node{{a.X, a.Y, 0}}
	visited[a.Grid.index(a.X, a.Y)] = true

	dirs := [4]struct{ dx, dy int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if a.Vision > 0 && n.d > a.Vision {
			continue
		}
		if n.d != 0 { // skip starting position
			state, _ := a.Grid.Get(n.x, n.y)
			if state == Food {
				return n.x, n.y, true
			}
		}
		for _, d := range dirs {
			nx, ny := n.x+d.dx, n.y+d.dy
			if !a.Grid.InBounds(nx, ny) {
				continue
			}
			idx := a.Grid.index(nx, ny)
			if visited[idx] {
				continue
			}
			s, _ := a.Grid.Get(nx, ny)
			if s == Obstacle {
				continue
			}
			visited[idx] = true
			q = append(q, node{nx, ny, n.d + 1})
		}
	}
	return 0, 0, false
}

// MoveTowards searches for food and moves one step toward it if found.
func (a *Agent) MoveTowardsFood() error {
	fx, fy, found := a.SearchFood()
	if !found {
		return nil
	}
	dx, dy := 0, 0
	if fx != a.X {
		dx = 1
		if fx < a.X {
			dx = -1
		}
	} else if fy != a.Y {
		dy = 1
		if fy < a.Y {
			dy = -1
		}
	}
	if dx != 0 {
		if err := a.Move(dx, 0); err == nil {
			return nil
		}
	}
	if dy != 0 {
		if err := a.Move(0, dy); err == nil {
			return nil
		}
	}
	return ErrBlocked
}
