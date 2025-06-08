package world

import (
	"math/rand"
	"sort"
)

type Genome struct {
	Vision   int
	MoveCost int
	FoodGain int
}

func randomGenome(r *rand.Rand) Genome {
	return Genome{
		Vision:   r.Intn(10) + 1,
		MoveCost: r.Intn(3) + 1,
		FoodGain: r.Intn(5) + 1,
	}
}

func (g Genome) newAgent(grid *Grid) *Agent {
	return &Agent{
		Grid:     grid,
		Energy:   20,
		Vision:   g.Vision,
		MoveCost: g.MoveCost,
		FoodGain: g.FoodGain,
	}
}

func evaluateGenome(r *rand.Rand, genome Genome) int {
	grid := NewGrid(5, 5)
	_ = grid.Set(4, 4, Food)
	agent := genome.newAgent(grid)
	sim := &Simulation{Grid: grid, Agents: []*Agent{agent}}
	for i := 0; i < 10; i++ {
		sim.Step()
	}
	dist := abs(agent.X-4) + abs(agent.Y-4)
	penalty := genome.Vision/4 + genome.FoodGain/4
	return agent.Energy - dist - penalty
}

func crossover(r *rand.Rand, a, b Genome) Genome {
	g := Genome{}
	if r.Intn(2) == 0 {
		g.Vision = a.Vision
	} else {
		g.Vision = b.Vision
	}
	if r.Intn(2) == 0 {
		g.MoveCost = a.MoveCost
	} else {
		g.MoveCost = b.MoveCost
	}
	if r.Intn(2) == 0 {
		g.FoodGain = a.FoodGain
	} else {
		g.FoodGain = b.FoodGain
	}
	return g
}

func mutate(r *rand.Rand, g Genome) Genome {
	if r.Intn(5) == 0 {
		g.Vision += r.Intn(3) - 1
		if g.Vision < 1 {
			g.Vision = 1
		}
		if g.Vision > 10 {
			g.Vision = 10
		}
	}
	if r.Intn(5) == 0 {
		g.MoveCost += r.Intn(3) - 1
		if g.MoveCost < 1 {
			g.MoveCost = 1
		}
		if g.MoveCost > 3 {
			g.MoveCost = 3
		}
	}
	if r.Intn(5) == 0 {
		g.FoodGain += r.Intn(3) - 1
		if g.FoodGain < 1 {
			g.FoodGain = 1
		}
		if g.FoodGain > 5 {
			g.FoodGain = 5
		}
	}
	return g
}

func Evolve(r *rand.Rand, generations, popSize int) Genome {
	if popSize < 2 {
		panic("population size must be at least 2")
	}
	pop := make([]Genome, popSize)
	for i := range pop {
		pop[i] = randomGenome(r)
	}
	best := pop[0]
	bestScore := evaluateGenome(r, best)
	for g := 0; g < generations; g++ {
		scores := make([]int, popSize)
		for i, gen := range pop {
			scores[i] = evaluateGenome(r, gen)
			if scores[i] > bestScore {
				bestScore = scores[i]
				best = gen
			}
		}
		sort.SliceStable(pop, func(i, j int) bool { return scores[i] > scores[j] })
		parents := pop[:2]
		newPop := make([]Genome, 0, popSize)
		newPop = append(newPop, parents...)
		for len(newPop) < popSize {
			child := mutate(r, crossover(r, parents[0], parents[1]))
			newPop = append(newPop, child)
		}
		pop = newPop
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
