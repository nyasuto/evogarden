package world

// Simulation manages a collection of agents within a grid.
type Simulation struct {
	Grid   *Grid
	Agents []*Agent
}

// Step advances the simulation by one tick. Each agent will attempt to move
// toward the nearest food source if it still has energy.
func (s *Simulation) Step() {
	for _, a := range s.Agents {
		if a.Energy <= 0 {
			continue
		}
		_ = a.MoveTowardsFood()
	}
}
